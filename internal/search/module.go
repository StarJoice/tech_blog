package search

import (
	"github.com/StarJoice/tech_blog/internal/search/internal/event"
	"github.com/StarJoice/tech_blog/internal/search/internal/service"
	"github.com/StarJoice/tech_blog/internal/search/internal/web"
)

type Module struct {
	Svc Service
	Hdl *Handler
	c   *event.SyncConsumer
}

type Service = service.SearchService
type Handler = web.Handler
