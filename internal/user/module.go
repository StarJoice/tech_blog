//@Date 2024/12/12 16:19
//@Desc

package user

import (
	"github.com/StarJoice/tech_blog/internal/user/internal/service"
	"github.com/StarJoice/tech_blog/internal/user/internal/web"
)

type Handler = web.UserHandler
type Service = service.UserService

type Module struct {
	Svc Service
	Hdl *Handler
}
