package codec

import (
	"crypto/md5"
	"fmt"
	"net/url"
	"sort"
	"strings"
)

func init() {
	RegisterSigner(SignerMd5, &md5sign{})
}

type md5sign struct{}

func (s *md5sign) Sign(kvs map[string]string, secret string) string {

	// 不参与签名的字段
	unsignField := map[string]struct{}{
		"c":        {},
		"apiName":  {},
		"sig":      {},
		"fromGame": {},
		"backUrl":  {},
		"a":        {},
	}

	temp := make(map[string]string, len(kvs))
	for k, v := range kvs {
		if _, ok := unsignField[k]; ok {
			continue
		}
		temp[k] = v
	}

	str := sortMap(temp)

	str += "+" + secret

	return fmt.Sprintf("%x", md5.Sum([]byte(str)))

}

func sortMap(kvs map[string]string) string {

	n := 0
	// 先给key按照字典序排列
	keys := make([]string, 0, len(kvs))
	for k, v := range kvs {
		keys = append(keys, k)
		n += len(v) + 1
	}
	sort.Strings(keys)

	builder := strings.Builder{}
	builder.Grow(n)
	builder.WriteString(url.QueryEscape(kvs[keys[0]])) // 每个值先进行encode,再参与签名
	for i := 1; i < len(keys); i++ {
		builder.WriteByte('+')
		builder.WriteString(url.QueryEscape(kvs[keys[i]]))
	}

	str := builder.String()

	return str
}
