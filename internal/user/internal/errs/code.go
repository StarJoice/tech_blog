package errs

var (
	SystemError = ErrorCode{
		Code: 501001,
		Msg:  "系统错误",
	}
	DataError = ErrorCode{
		Code: 501002,
		Msg:  "数据格式出错",
	}
)

type ErrorCode struct {
	Code int
	Msg  string
}
