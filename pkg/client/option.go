package client

import (
	"fmt"
	"net/http"

	"github.com/tglivelink/livelink-go/pkg/codec"
)

type Option struct {
	Serializer codec.SerializerType // 序列化方式
	Coder      codec.CodeType       // 信息加解密方式
	Signer     codec.SignType       // 签名方式

	// 自定义http请求客户端
	HttpClient interface {
		Do(*http.Request) (*http.Response, error)
	}
}

func NewOption() *Option {
	return &Option{
		Serializer: codec.SerializerJson,
		Coder:      codec.CoderEcb,
		Signer:     codec.SignerMd5,
		HttpClient: nil,
	}
}

func (o *Option) Check() error {
	if codec.GetCoder(o.Coder) == nil {
		return fmt.Errorf("coder:%d is unsupported", o.Coder)
	}
	if codec.GetSigner(o.Signer) == nil {
		return fmt.Errorf("signer:%d is unsupported", o.Signer)
	}
	if codec.GetSerializer(o.Serializer) == nil {
		return fmt.Errorf("serializer:%d is unsupported", o.Serializer)
	}
	return nil
}

func (o *Option) Clone() *Option {
	t := *o
	return &t
}

type Options func(o *Option)

func WithSigner(t codec.SignType) Options {
	return func(o *Option) {
		o.Signer = t
	}
}
