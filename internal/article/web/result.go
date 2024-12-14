package web

import (
	"github.com/StarJoice/tech_blog/internal/article/errs"
	ginx "github.com/StarJoice/tools/ginx/wrapper"
)

var (
	systemErrorResult = ginx.Result{
		Code: errs.SystemError.Code,
		Msg:  errs.SystemError.Msg,
	}
)
