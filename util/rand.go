package util

import (
	"crypto/rand"
	"io"
)

// RandBytes 随机返回字节，长度为n
func RandBytes(n int) []byte {
	dst := make([]byte, n)
	io.ReadFull(rand.Reader, dst)
	return dst
}
