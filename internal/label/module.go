package label

import (
	"github.com/StarJoice/tech_blog/internal/label/service"
	"github.com/StarJoice/tech_blog/internal/label/web"
)

type Module struct {
	Hdl *web.Handler
}
type Handler = web.Handler
type Service = service.Service
