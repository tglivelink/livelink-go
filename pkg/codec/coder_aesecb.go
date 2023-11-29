package codec

import (
	"crypto/aes"
	"encoding/base64"
)

// ecb模式安全性不足，go原生没有提供，只能自己写

func init() {
	RegisterCoder(CoderEcb, &aesEcb{})
}

type aesEcb struct{}

func (a *aesEcb) Encrypt(data []byte, secret string) ([]byte, error) {
	block, err := aes.NewCipher([]byte(secret))
	if err != nil {
		return nil, err
	}

	// 填充，遵循PKCS7规则
	padding := block.BlockSize() - len(data)%block.BlockSize()
	tmp := make([]byte, 0, len(data)+padding)
	tmp = append(tmp, data...)
	for i := 0; i < padding; i++ {
		tmp = append(tmp, byte(padding))
	}

	result := make([]byte, len(tmp))

	// 按组分别加密
	for i, size := 0, block.BlockSize(); i < len(result); i += size {
		block.Encrypt(result[i:i+size], tmp[i:i+size])
	}

	dst := make([]byte, base64.StdEncoding.EncodedLen(len(result)))
	base64.StdEncoding.Encode(dst, result)
	return dst, nil
}

func (a *aesEcb) Decrypt(data []byte, secret string) ([]byte, error) {
	bs := make([]byte, base64.StdEncoding.DecodedLen(len(data)))
	n, err := base64.StdEncoding.Decode(bs, data)
	if err != nil {
		return nil, err
	}
	bs = bs[:n]

	block, err := aes.NewCipher([]byte(secret))
	if err != nil {
		return nil, err
	}

	result := make([]byte, len(bs))

	// 按组分别解密
	for i, size := 0, block.BlockSize(); i < len(result); i += size {
		block.Decrypt(result[i:i+size], bs[i:i+size])
	}

	// 去掉填充，遵循PKCS7规则
	padding := result[len(result)-1]
	result = result[:len(result)-int(padding)]

	return result, nil
}
