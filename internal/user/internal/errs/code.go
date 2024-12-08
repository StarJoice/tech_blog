package errs

var (
	SystemError = ErrorCode{
		Code: 101001,
		Msg:  "系统错误",
	}
	DataError = ErrorCode{
		Code: 101002,
		Msg:  "数据格式出错",
	}
)

type ErrorCode struct {
	Code int
	Msg  string
}
