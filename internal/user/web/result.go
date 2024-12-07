package web

import (
	"github.com/StarJoice/tech_blog/internal/user/errs"
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

type Result struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}
