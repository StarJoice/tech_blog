package interactive

import (
	"github.com/StarJoice/tech_blog/internal/interactive/internal/event"
)

type Module struct {
	Hdl *Handler
	Svc Service
	c   *event.Consumer
}
