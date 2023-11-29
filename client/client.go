package client

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/huangzixiang5/livelink-go/codec"
	"github.com/huangzixiang5/livelink-go/config"
	"github.com/huangzixiang5/livelink-go/log"
	"github.com/huangzixiang5/livelink-go/util"
)

// Client 客户端接口
type Client interface {
	Do(ctx context.Context, head *ReqHead, req interface{}, rsp interface{}, opts ...Options) error
}

var DefaultClient Client = &client{}

//////////////////////////////////////

// 默认使用http客户端
type client struct {
	opt     *Option
	optOnce sync.Once
}

func (c *client) Do(ctx context.Context, head *ReqHead, req interface{},
	rsp interface{}, opts ...Options) error {

	ctx = util.WithTraceID(ctx, fmt.Sprintf("%x", util.RandBytes(8)))

	opt := c.getOption(opts...)
	if err := c.checkOpt(opt); err != nil {
		return err
	}

	if err := c.checkHead(head, opt); err != nil {
		return err
	}

	kvs, err := c.getSignMap(head, req, opt)
	if err != nil {
		return err
	}

	httpReq, err := c.getHttpReq(ctx, head, kvs, req, opt)
	if err != nil {
		return err
	}

	httpRsp, err := opt.HttpClient.Do(httpReq)
	if err != nil {
		return err
	}

	if err := c.parserHttpRsp(ctx, httpRsp, rsp, opt); err != nil {
		return err
	}

	return nil
}

func (c *client) getOption(opts ...Options) *Option {
	c.optOnce.Do(func() {
		cfg := config.GlobalConfig()
		c.opt = &Option{
			Serializer: cfg.Serializer,
			Coder:      cfg.Coder,
			Signer:     cfg.Signer,
			HttpClient: httpClient,
		}
	})

	var opt *Option

	if len(opts) == 0 {
		opt = c.opt
	} else {
		opt = c.opt.Clone()
		for _, v := range opts {
			v(opt)
		}
	}

	return opt
}

func (c *client) checkOpt(opt *Option) error {
	return opt.Check()
}

func (c *client) checkHead(head *ReqHead, opt *Option) error {
	fromGame := head.FromGame
	if head.LivePlatId == "" && !fromGame {
		head.LivePlatId = config.GlobalConfig().Appid
	}
	if head.GameId == "" && fromGame {
		head.GameId = config.GlobalConfig().Appid
	}

	return nil
}

func (c *client) getSignMap(head *ReqHead, req interface{}, opt *Option) (map[string]string, error) {

	kvs := make(map[string]string, 16)

	kvs["livePlatId"] = head.LivePlatId
	kvs["actId"] = fmt.Sprintf("%d", head.ActId)
	kvs["gameId"] = head.GameId
	kvs["t"] = fmt.Sprintf("%d", time.Now().Unix())
	kvs["nonce"] = fmt.Sprintf("%x", util.RandBytes(3))
	if head.FromGame {
		kvs["fromGame"] = "1"
	}
	if head.PathOrApiName != "" && head.PathOrApiName[0] != '/' {
		kvs["apiName"] = head.PathOrApiName
	}
	for k, v := range head.Ext {
		kvs[k] = v
	}

	user, err := codec.GetSerializer(opt.Serializer).Marshal(head.User)
	if err != nil {
		return nil, fmt.Errorf("error occurred when serialize user: %w", err)
	}
	user2, err := codec.GetCoder(opt.Coder).Encrypt(user, config.GlobalConfig().SecKey)
	if err != nil {
		return nil, fmt.Errorf("error occurred when encrypt user: %w", err)
	}
	kvs["code"] = string(user2)

	kvs["sig"] = codec.GetSigner(opt.Signer).Sign(kvs, config.GlobalConfig().SigKey)

	return kvs, nil
}

func (c *client) getHttpReq(ctx context.Context, head *ReqHead,
	kvs map[string]string, req interface{}, opt *Option) (*http.Request, error) {

	query := url.Values{}
	for k, v := range kvs {
		query.Add(k, v)
	}

	addr := c.mergePath(c.getDomain(opt), head.PathOrApiName)
	addr += "?" + query.Encode()

	log.Infof(ctx, "request:%s", addr)

	var bs []byte
	var err error

	if req != nil {
		if bs, err = codec.GetSerializer(opt.Serializer).Marshal(req); err != nil {
			return nil, err
		}
	} else {
		bs = []byte(`{}`)
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, addr, bytes.NewBuffer(bs))
	if err != nil {
		return nil, fmt.Errorf("error occurred when create http.req: %w", err)
	}

	return httpReq, nil
}

func (c *client) getDomain(opt *Option) string {
	return config.GlobalConfig().Domain
}

func (c *client) mergePath(domain, path string) string {
	if path == "" {
		return domain
	}

	if path[0] != '/' { // 旧版的apiName形式
		return domain + "/livelink"
	}

	addr := ""

	a, b := domain[len(domain)-1] == '/', path[0] == '/'

	switch {
	case a && b:
		addr = domain + path[1:]
	case a || b:
		addr = domain + path
	case !a && !b:
		addr = domain + "/" + path
	}
	return addr
}

func (c *client) parserHttpRsp(ctx context.Context, httpRsp *http.Response, rsp interface{}, opt *Option) error {
	defer httpRsp.Body.Close()

	if httpRsp.StatusCode != http.StatusOK {
		return fmt.Errorf("http.status error [%d,%s]", httpRsp.StatusCode, httpRsp.Status)
	}

	bs, err := io.ReadAll(httpRsp.Body)
	if err != nil {
		return fmt.Errorf("error occurred when read http.response: %w", err)
	}

	log.Infof(ctx, "response:%s", bs)

	switch v := rsp.(type) {
	case nil:
		return nil
	case *[]byte:
		*v = bs
	default:
		if err := codec.GetSerializer(opt.Serializer).Unmarshal(bs, rsp); err != nil {
			return fmt.Errorf("error occurred when serialize response: %w", err)
		}
	}
	return nil
}

var httpClient = &http.Client{
	Transport: &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		ForceAttemptHTTP2:     true,
		IdleConnTimeout:       60 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		MaxIdleConnsPerHost:   100,
		DisableCompression:    true,
		ExpectContinueTimeout: time.Second,
	},
	Jar:     nil,
	Timeout: time.Second * 30,
}
