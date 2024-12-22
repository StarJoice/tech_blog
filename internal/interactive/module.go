package interactive

import (
	"github.com/StarJoice/tech_blog/internal/interactive/service"
	"github.com/StarJoice/tech_blog/internal/interactive/web"
)

type Module struct {
	Hdl *Handler
	Svc Service
}

type Handler = web.Handler
type Service = service.Service
