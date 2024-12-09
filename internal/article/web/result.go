//@Date 2024/12/9 21:19
//@Desc

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
