package client

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/tglivelink/livelink-go/pkg/codec"
	"github.com/tglivelink/livelink-go/pkg/log"
)

// Client 客户端接口
type Client interface {
	Do(ctx context.Context, head *Head, opts ...Options) error
}

var DefaultClient Client = New(Secret{})

func New(secret Secret) Client {
	opt := NewOption()
	opt.SecKey = secret.SecKey
	opt.SigKey = secret.SigKey
	return &client{opt: opt}
}

/***************************************/

// 默认使用http客户端
type client struct {
	opt *Option
}

func (c *client) Do(ctx context.Context, head *Head, opts ...Options) (err error) {

	opt := c.getOption(opts...)
	if err := c.checkOpt(opt); err != nil {
		return err
	}

	if err := c.doWithOption(ctx, head, opt); err != nil {
		return err
	}

	return nil
}

func (c *client) getOption(opts ...Options) *Option {

	if len(opts) == 0 {
		return c.opt
	}

	opt := c.opt.Clone()
	for _, v := range opts {
		v(opt)
	}
	return opt
}

func (c *client) checkOpt(opt *Option) error {
	if err := opt.Check(); err != nil {
		return err
	}
	return nil
}

func (c *client) doWithOption(ctx context.Context, head *Head, opt *Option) error {

	if opt.Timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, opt.Timeout)
		defer cancel()
	}

	if err := c.checkHead(head, opt); err != nil {
		return err
	}

	kvs, err := c.getSignedMap(head, opt)
	if err != nil {
		return err
	}

	httpReq, err := c.getHttpReq(ctx, head, kvs, opt)
	if err != nil {
		return err
	}

	httpRsp, err := opt.HttpClient.Do(httpReq)
	if err != nil {
		return err
	}

	if err := c.parserHttpRsp(ctx, httpRsp, head.Rsp, opt); err != nil {
		return err
	}
	return nil
}

func (c *client) checkHead(head *Head, opt *Option) error {
	if head.PathOrApiName == "" {
		return fmt.Errorf("请求路径不能为空")
	}
	return nil
}

func (c *client) getSignedMap(head *Head, opt *Option) (map[string]string, error) {

	kvs := head.Param.FixToKVs()
	if head.PathOrApiName != "" && head.PathOrApiName[0] != '/' {
		kvs["apiName"] = head.PathOrApiName
	}
	for k, v := range head.Param.Ext {
		kvs[k] = v
	}

	user, err := codec.GetSerializer(opt.Serializer).Marshal(head.Param.User)
	if err != nil {
		return nil, fmt.Errorf("error occurred when serialize user: %w", err)
	}
	user2, err := codec.GetCoder(opt.Coder).Encrypt(user, opt.SecKey)
	if err != nil {
		return nil, fmt.Errorf("error occurred when encrypt user: %w", err)
	}
	kvs["code"] = string(user2)

	kvs["sig"] = codec.GetSigner(opt.Signer).Sign(kvs, opt.SigKey)

	return kvs, nil
}

func (c *client) getHttpReq(ctx context.Context, head *Head, kvs map[string]string, opt *Option) (*http.Request, error) {

	query := make(url.Values, len(kvs))
	for k, v := range kvs {
		query.Add(k, v)
	}

	addr := c.mergePath(opt.Domain, head.PathOrApiName)
	addr += "?" + query.Encode()

	log.InfoContextf(ctx, "request:%s", addr)

	var bs []byte
	var err error

	if head.Body != nil {
		if bs, err = codec.GetSerializer(opt.Serializer).Marshal(head.Body); err != nil {
			return nil, err
		}
	} else {
		bs = []byte(`{}`)
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, addr, bytes.NewBuffer(bs))
	if err != nil {
		return nil, fmt.Errorf("error occurred when create http.req: %w", err)
	}

	switch opt.Serializer {
	case codec.SerializerJson:
		httpReq.Header.Set("Content-Type", "application/json")
	}

	return httpReq, nil
}

func (c *client) mergePath(domain, path string) string {
	if path == "" {
		return domain
	}

	if path[0] != '/' { // 旧版的apiName形式,固定前缀是livelink
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

	log.InfoContextf(ctx, "response:%s", bs)

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

func httpClient() *http.Client {
	return &http.Client{
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
			MaxIdleConnsPerHost:   300,
			DisableCompression:    true,
			ExpectContinueTimeout: time.Second,
		},
		Jar:     nil,
		Timeout: time.Second * 60,
	}
}
