package web

import (
	"github.com/StarJoice/tech_blog/internal/user/internal/errs"
	ginx "github.com/StarJoice/tools/ginx/wrapper"
)

var (
	systemErrorResult = ginx.Result{
		Code: errs.SystemError.Code,
		Msg:  errs.SystemError.Msg,
	}
	DataErrorResult = ginx.Result{
		Code: errs.DataError.Code,
		Msg:  errs.DataError.Msg,
	}
)

// Result 返回的结果结构体
// @Description 该结构体描述了 API 请求的返回格式
// @object
type Result struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}
