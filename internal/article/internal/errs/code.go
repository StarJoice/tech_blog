package errs

// article 模块使用502标识

var (
	SystemError = ErrorCode{Code: 502001, Msg: "系统错误"}
)

type ErrorCode struct {
	Code int
	Msg  string
}
