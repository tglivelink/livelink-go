package codec

import (
	"crypto/md5"
	"fmt"
)

func init() {
	RegisterSigner(SignerMd5Web, &md5WebSign{})
}

type md5WebSign struct{}

func (h *md5WebSign) Sign(kvs map[string]string, secret string) string {

	// 最多只需要这几个字段参与签名
	fields := []string{
		"livePlatId", "gameIdList", "t", "code", "gameAuthScene",
	}

	temp := make(map[string]string, len(fields))
	for _, key := range fields {
		value, ok := kvs[key]
		if ok {
			temp[key] = value
		}
	}

	str := sortMap(temp)

	str += "+" + secret

	return fmt.Sprintf("%x", md5.Sum([]byte(str)))

}
