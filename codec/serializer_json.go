package codec

import "encoding/json"

func init() {
	RegisterSerializer("", &sJson{})
	RegisterSerializer(SerializerJson, &sJson{})
}

type sJson struct{}

func (*sJson) Marshal(data interface{}) ([]byte, error) {
	return json.Marshal(data)
}

func (*sJson) Unmarshal(bs []byte, data interface{}) error {
	return json.Unmarshal(bs, data)
}
