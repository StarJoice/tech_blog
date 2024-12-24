package web

import (
	"github.com/StarJoice/tech_blog/internal/comment/internal/service"
	ginx "github.com/StarJoice/tools/ginx/wrapper"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	svc service.Service
}

func NewHandler(svc service.Service) *Handler {
	return &Handler{svc: svc}
}

// PrivateRoutes 评论模块要登陆后才能使用
func (h *Handler) PrivateRoutes(server *gin.Engine) {
	server.POST("/comment", ginx.WithRequest[createCommentReq](h.CreateComment))
}

func (h *Handler) PublicRoutes(server *gin.Engine) {
	//TODO implement me
	panic("implement me")
}

type createCommentReq struct {
}

func (h *Handler) CreateComment(ctx *ginx.Context,
	req createCommentReq) (ginx.Result, error) {

}
