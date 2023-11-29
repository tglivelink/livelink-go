package codec

type SignType int

// 目前支持的签名方式
const (
	SignerMd5      SignType = 0
	SignerMd5Fixed SignType = 1
)

type CodeType int

// 目前支持的用户信息加解密方式
const (
	CoderEcb CodeType = 0
)

type SerializerType int

// 数据序列化方式
const (
	SerializerJson SerializerType = 0
)

// Serializer 序列化/反序列化接口
type Serializer interface {
	Marshal(data interface{}) ([]byte, error)
	Unmarshal(bs []byte, data interface{}) error
}

// Singer 签名接口
type Singer interface {
	Sign(kvs map[string]string, secret string) string
}

// Coder 敏感信息加解密接口
type Coder interface {
	Encrypt(text []byte, secret string) ([]byte, error)
	Decrypt(text []byte, secret string) ([]byte, error)
}

var serializers = map[SerializerType]Serializer{}
var singers = map[SignType]Singer{}
var coders = map[CodeType]Coder{}

func RegisterSerializer(t SerializerType, s Serializer) {
	serializers[t] = s
}

func GetSerializer(t SerializerType) Serializer {
	return serializers[t]
}

func RegisterSigner(t SignType, s Singer) {
	singers[t] = s
}

func GetSigner(t SignType) Singer {
	return singers[t]
}

func RegisterCoder(t CodeType, c Coder) {
	coders[t] = c
}

func GetCoder(t CodeType) Coder {
	return coders[t]
}
