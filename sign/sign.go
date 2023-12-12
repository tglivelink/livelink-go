package sign

import (
	"fmt"
	"net/url"

	"github.com/tglivelink/livelink-go/pkg/client"
	"github.com/tglivelink/livelink-go/pkg/codec"
)

/**************************/

type MiniProgramReq struct {
	client.Param
	FaceUrl  string `json:"faceUrl"`  // 拉起小程序展示的头像，必填
	NickName string `json:"nickName"` // 拉起小程序展示的昵称，必填
}

// SignForMiniProgram 生成小程序签名
func SignForMiniProgram(param *MiniProgramReq, secret *client.Secret, opts ...client.Options) (url.Values, error) {

	if param.FaceUrl == "" {
		return nil, fmt.Errorf("缺少faceUrl参数")
	}
	if param.NickName == "" {
		return nil, fmt.Errorf("缺少nickName参数")
	}

	param.AddExt("faceUrl", param.FaceUrl)
	param.AddExt("nickName", param.NickName)

	opts = append([]client.Options{client.WithSigner(codec.SignerMd5Fixed)}, opts...)
	return Sign(&param.Param, secret, opts...)
}

// Sign 计算签名，将返回拼接在url后发起请求即可
func Sign(param *client.Param, secret *client.Secret, opts ...client.Options) (url.Values, error) {

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
	code, err := codec.GetCoder(opt.Coder).Encrypt(userByte, secret.SecKey)
	if err != nil {
		return nil, err
	}

	kvs := param.FixToKVs()
	kvs["code"] = string(code)

	sig := codec.GetSigner(opt.Signer).Sign(kvs, secret.SigKey)
	kvs["sig"] = sig

	t := make(url.Values, len(kvs))
	for k, v := range kvs {
		t.Add(k, v)
	}

	return t, nil
}

/************************************/
