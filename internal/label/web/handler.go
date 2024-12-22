package web

import (
	"github.com/StarJoice/tech_blog/internal/label/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	svc service.Service
}

func NewHandler(svc service.Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) PublicRoutes(server *gin.Engine) {
	server.POST("/label")
}
