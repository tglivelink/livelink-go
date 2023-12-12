package errs

type RetCode int

const (
	RetCodeSucc RetCode = 0

	// RetCodeParamErr 参数错误，参照msg返回提示
	RetCodeParamErr RetCode = -999999
	// RetCodeSigErr 签名错误
	RetCodeSigErr RetCode = -1000
	// RetCodeExpiredErr 请求过期
	RetCodeExpiredErr RetCode = -4003
)
