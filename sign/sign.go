package sign

import (
	"net/url"

	"github.com/huangzixiang5/livelink-go/pkg/client"
	"github.com/huangzixiang5/livelink-go/pkg/codec"
)

/**************************/

// Sign 计算签名，将返回拼接在url后发起请求即可
func Sign(param *client.ReqParam, sigKey, secKey string, opts ...client.Options) (url.Values, error) {

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
	code, err := codec.GetCoder(opt.Coder).Encrypt(userByte, secKey)
	if err != nil {
		return nil, err
	}

	kvs := param.FixToKVs()
	kvs["code"] = string(code)

	sig := codec.GetSigner(opt.Signer).Sign(kvs, sigKey)
	kvs["sig"] = sig

	t := url.Values{}
	for k, v := range kvs {
		t.Add(k, v)
	}

	return t, nil
}

/************************************/
