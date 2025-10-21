package sign

import (
	"fmt"
	"net/url"
	"time"

	"github.com/tglivelink/livelink-go/pkg/client"
	"github.com/tglivelink/livelink-go/pkg/codec"
	"github.com/tglivelink/livelink-go/pkg/util"
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

type LivelinkLogin struct {
	T     int64
	Nonce string
	User  *client.GameUser
}

// SignForLivelinkLoginType 游戏反绑，使用LiveLink登录模式
// see also https://livelink.qq.com/doc/activities/pages/bind/web-sdk/lm-livelink.html
func SignForLivelinkLoginType(param *LivelinkLogin, secret *client.Secret, opts ...client.Options) (code, sign string, err error) {

	if param.User == nil || param.User.Key() == "" {
		return "", "", fmt.Errorf("缺少User参数")
	}

	if param.T <= 0 {
		param.T = time.Now().Unix()
	}
	if param.Nonce == "" {
		param.Nonce = fmt.Sprintf("%x", util.RandBytes(3))
	}

	opt := client.Option{}
	for _, v := range opts {
		v(&opt)
	}

	if err = opt.Check(); err != nil {
		return
	}

	var userByte []byte
	userByte, err = codec.GetSerializer(opt.Serializer).Marshal(param.User)
	if err != nil {
		return
	}
	userByte, err = codec.GetCoder(opt.Coder).Encrypt(userByte, secret.SecKey)
	if err != nil {
		return
	}
	code = string(userByte)

	sign = codec.GetSigner(opt.Signer).Sign(map[string]string{
		"t":     fmt.Sprintf("%d", param.T),
		"nonce": param.Nonce,
		"code":  code,
	}, secret.SigKey)

	return
}
