package client

import (
	"fmt"
	"net/http"

	"github.com/huangzixiang5/livelink-go/codec"
)

type Option struct {
	Serializer string // 序列化方式
	Coder      string // 信息加解密方式
	Signer     string // 签名方式

	// 自定义http请求客户端
	HttpClient interface {
		Do(*http.Request) (*http.Response, error)
	}
}

func (o *Option) Check() error {
	if codec.GetCoder(o.Coder) == nil {
		return fmt.Errorf("coder:%s is unsupported", o.Coder)
	}
	if codec.GetSigner(o.Signer) == nil {
		return fmt.Errorf("signer:%s is unsupported", o.Signer)
	}
	if codec.GetSerializer(o.Serializer) == nil {
		return fmt.Errorf("serializer:%s is unsupported", o.Serializer)
	}
	return nil
}

func (o *Option) Clone() *Option {
	t := *o
	return &t
}

type Options func(o *Option)

func WithSigner(s string) Options {
	return func(o *Option) {
		o.Signer = s
	}
}
