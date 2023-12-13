package client

import (
	"fmt"
	"net/http"
	"time"

	"github.com/tglivelink/livelink-go/pkg/codec"
)

var (
	Domain = "https://s1.livelink.qq.com"
)

type Option struct {
	Serializer codec.SerializerType
	Coder      codec.CodeType
	Signer     codec.SignType

	Domain string
	SigKey string
	SecKey string

	Timeout time.Duration

	HttpClient interface {
		Do(*http.Request) (*http.Response, error)
	}
}

func NewOption() *Option {
	opt := &Option{
		Serializer: codec.SerializerJson,
		Coder:      codec.CoderEcb,
		Signer:     codec.SignerMd5,
		Domain:     Domain,
		HttpClient: httpClient(),
	}
	return opt
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

/**********/

// WithSigner
func WithSigner(t codec.SignType) Options {
	return func(o *Option) {
		o.Signer = t
	}
}

// WithCoder
func WithCoder(c codec.CodeType) Options {
	return func(o *Option) {
		o.Coder = c
	}
}

// WithHttpClient
func WithHttpClient(c interface {
	Do(*http.Request) (*http.Response, error)
}) Options {
	return func(o *Option) {
		o.HttpClient = c
	}
}

// WithSecret
func WithSecret(s Secret) Options {
	return func(o *Option) {
		o.SecKey = s.SecKey
		o.SigKey = s.SigKey
	}
}

// WithDomain
func WithDomain(s string) Options {
	return func(o *Option) {
		o.Domain = s
	}
}

// WithTimeout
func WithTimeout(d time.Duration) Options {
	return func(o *Option) {
		o.Timeout = d
	}
}
