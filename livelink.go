package livelink

import (
	"fmt"
	"net/url"

	"github.com/huangzixiang5/livelink-go/client"
	"github.com/huangzixiang5/livelink-go/codec"
)

/**************************/

// MiniProgramReq 拉起小程序参数
type MiniProgramReq struct {
	GameIdList string          // 游戏列表，必填，类似： cf,lol
	LivePlatId string          // 平台id，必填
	T          int64           // 时间戳（秒），必填
	User       client.PlatUser // 用户uid，必填
	FaceUrl    string          // 用户所在平台的头像，必填
	NickName   string          // 用户所在平台昵称，必填

	Ext map[string]string // 扩展参数,可以参考文档添加
}

// ArgsForMiniProgram 生成拉起小程序需要的参数
func ArgsForMiniProgram(param *MiniProgramReq, sigKey, secKey string, opts ...client.Options) (string, error) {

	if param.GameIdList == "" || param.LivePlatId == "" ||
		param.T <= 0 || param.User.Userid == "" ||
		param.FaceUrl == "" || param.NickName == "" {
		return "", fmt.Errorf("必填参数不能为空")
	}

	opt := client.Option{
		Signer: codec.SignerMd5Web, // 默认使用这个签名
	}
	for _, v := range opts {
		v(&opt)
	}

	if err := opt.Check(); err != nil {
		return "", err
	}

	userByte, err := codec.GetSerializer(opt.Serializer).Marshal(param.User)
	if err != nil {
		return "", err
	}
	code, err := codec.GetCoder(opt.Coder).Encrypt(userByte, secKey)
	if err != nil {
		return "", err
	}

	kvs := map[string]string{
		"gameIdList": param.GameIdList,
		"livePlatId": param.LivePlatId,
		"code":       string(code),
		"t":          fmt.Sprintf("%d", param.T),
		"faceUrl":    param.FaceUrl,
		"nickName":   param.NickName,
	}
	for k, v := range param.Ext {
		kvs[k] = v
	}

	sig := codec.GetSigner(opt.Signer).Sign(kvs, sigKey)
	kvs["sig"] = sig

	t := url.Values{}
	for k, v := range kvs {
		t.Add(k, v)
	}

	return t.Encode(), nil
}

/************************************/

type IframeReq struct {
	GameId     string          // 游戏id,必填
	Timestamp  int64           // 时间戳（秒），必填
	LivePlatId string          // 平台id,必填
	ActId      uint64          // 活动id，必填
	User       client.PlatUser // 用户信息，必填
}

// ArgsForIframe 生成拉起内嵌活动iframe需要的参数
func ArgsForIframe(param *IframeReq, sigKey, secKey string, opts ...client.Options) (string, error) {

	if param.GameId == "" || param.LivePlatId == "" ||
		param.Timestamp <= 0 || param.User.Userid == "" || param.ActId <= 0 {
		return "", fmt.Errorf("必填参数不能为空")
	}

	opt := client.Option{}
	for _, v := range opts {
		v(&opt)
	}

	if err := opt.Check(); err != nil {
		return "", err
	}

	userByte, err := codec.GetSerializer(opt.Serializer).Marshal(param.User)
	if err != nil {
		return "", err
	}
	code, err := codec.GetCoder(opt.Coder).Encrypt(userByte, secKey)
	if err != nil {
		return "", err
	}

	kvs := map[string]string{
		"gameId":     param.GameId,
		"livePlatId": param.LivePlatId,
		"code":       string(code),
		"timestamp":  fmt.Sprintf("%d", param.Timestamp),
		"v":          "2.0",
		"actId":      fmt.Sprintf("%d", param.ActId),
		"appId":      param.LivePlatId,
	}

	sig := codec.GetSigner(opt.Signer).Sign(kvs, sigKey)
	kvs["sig"] = sig

	t := url.Values{}
	for k, v := range kvs {
		t.Add(k, v)
	}

	return t.Encode(), nil
}
