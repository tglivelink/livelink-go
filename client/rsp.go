package client

type RspHead struct {
	RspBase
	Data interface{} `json:"jData"`
}

type RspBase struct {
	Ret     int    `json:"iRet"`
	Msg     string `json:"sMsg"`
	Tid     string `json:"tid"`
	ApiName string `json:"apiName"`
}
