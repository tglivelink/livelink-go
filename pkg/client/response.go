package client

import (
	"strings"

	"github.com/tglivelink/livelink-go/pkg/errs"
)

type Response struct {
	ResponseBase
	Data interface{} `json:"jData"`
}

type ResponseBase struct {
	Ret     errs.RetCode `json:"iRet"`
	Msg     string       `json:"sMsg"`
	Tid     string       `json:"tid"`
	ApiName string       `json:"apiName"`
}

func (rb *ResponseBase) GetCode() errs.RetCode {
	return rb.Ret
}

func (r *Response) Value(path string) interface{} {
	if path == "" {
		return r.Data
	}
	v := r.Data
	for _, key := range strings.Split(path, ".") {
		data, ok := v.(map[string]interface{})
		if ok {
			v = data[key]
		} else {
			return nil
		}
	}

	return v
}
