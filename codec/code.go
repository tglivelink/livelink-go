package codec

// 目前支持的签名方式
const (
	SignerMd5    = "md5"
	SignerMd5Web = "md5_web"
)

// 目前支持的用户信息加解密方式
const (
	CoderEcb = "ecb"
)

// 数据序列化方式
const (
	SerializerJson = "json"
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

var serializers = map[string]Serializer{}
var singers = map[string]Singer{}
var coders = map[string]Coder{}

func RegisterSerializer(name string, s Serializer) {
	serializers[name] = s
}

func GetSerializer(name string) Serializer {
	return serializers[name]
}

func RegisterSigner(name string, s Singer) {
	singers[name] = s
}

func GetSigner(name string) Singer {
	return singers[name]
}

func RegisterCoder(name string, c Coder) {
	coders[name] = c
}

func GetCoder(name string) Coder {
	return coders[name]
}
