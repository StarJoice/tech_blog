package web

import (
	"github.com/StarJoice/tech_blog/internal/search/internal/service"
	ginx "github.com/StarJoice/tools/ginx/wrapper"
	"github.com/gin-gonic/gin"
	"github.com/gotomicro/ego/core/elog"
)

type Handler struct {
	svc    service.SearchService
	logger *elog.Component
}

func NewHandler(svc service.SearchService) *Handler {
	return &Handler{svc: svc, logger: elog.DefaultLogger}
}

func (h *Handler) PrivateRoutes(server *gin.Engine) {
	server.POST("/search", ginx.WithRequest[searchReq](h.Search))
}

func (h *Handler) PublicRoutes(server *gin.Engine) {
	//TODO implement me
	panic("implement me")
}

func (h *Handler) Search(ctx *ginx.Context, req searchReq) (ginx.Result, error) {
	data, err := h.svc.Search(ctx.Request.Context(), req.Offset, req.Limit, req.Keywords)
	if err != nil {
		return systemErrorResult, err
	}
	return ginx.Result{
		Msg:  "搜索成功",
		Data: data.Articles,
	}, nil
}
