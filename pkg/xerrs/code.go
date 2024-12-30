package xerrs

func NewSystemError(code int) ErrorCode {
	return ErrorCode{
		Code: code,
		Msg:  "系统错误",
	}
}
