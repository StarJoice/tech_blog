package label

import (
	"github.com/StarJoice/tech_blog/internal/label/internal/service"
	"github.com/StarJoice/tech_blog/internal/label/internal/web"
)

type Module struct {
	Hdl *web.Handler
}
type Handler = web.Handler
type Service = service.Service
