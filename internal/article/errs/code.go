//@Date 2024/12/9 20:35
//@Desc

package errs

var (
	SystemError = ErrorCode{Code: 502001, Msg: "系统错误"}
)

type ErrorCode struct {
	Code int
	Msg  string
}
