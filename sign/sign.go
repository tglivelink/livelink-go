package sign

import (
	"fmt"
	"net/url"

	"github.com/tglivelink/livelink-go/pkg/client"
	"github.com/tglivelink/livelink-go/pkg/codec"
	"github.com/tglivelink/livelink-go/pkg/config"
)

/**************************/

type MiniProgramReq struct {
	client.ReqParam
	FaceUrl  string `json:"faceUrl"`  // 拉起小程序展示的头像，必填
	NickName string `json:"nickName"` // 拉起小程序展示的昵称，必填
}

// SignForMiniProgram 生成小程序签名
func SignForMiniProgram(param *MiniProgramReq, cfg *config.ClientConfig, opts ...client.Options) (url.Values, error) {

	if param.FaceUrl == "" {
		return nil, fmt.Errorf("缺少faceUrl参数")
	}
	if param.NickName == "" {
		return nil, fmt.Errorf("缺少nickName参数")
	}
	if param.Ext == nil {
		param.Ext = map[string]string{}
	}
	param.Ext["faceUrl"] = param.FaceUrl
	param.Ext["nickName"] = param.NickName

	opts = append([]client.Options{client.WithSigner(codec.SignerMd5Fixed)}, opts...)
	return Sign(&param.ReqParam, cfg, opts...)
}

// Sign 计算签名，将返回拼接在url后发起请求即可
func Sign(param *client.ReqParam, cfg *config.ClientConfig, opts ...client.Options) (url.Values, error) {

	if param.LivePlatId == "" {
		return nil, fmt.Errorf("缺少LivePlatId参数")
	}
	if param.GameId == "" {
		return nil, fmt.Errorf("缺少GameId参数")
	}

	opt := client.Option{}
	for _, v := range opts {
		v(&opt)
	}

	if err := opt.Check(); err != nil {
		return nil, err
	}

	userByte, err := codec.GetSerializer(opt.Serializer).Marshal(param.User)
	if err != nil {
		return nil, err
	}
	code, err := codec.GetCoder(opt.Coder).Encrypt(userByte, cfg.SecKey)
	if err != nil {
		return nil, err
	}

	kvs := param.FixToKVs()
	kvs["code"] = string(code)

	sig := codec.GetSigner(opt.Signer).Sign(kvs, cfg.SigKey)
	kvs["sig"] = sig

	t := url.Values{}
	for k, v := range kvs {
		t.Add(k, v)
	}

	return t, nil
}

/************************************/
