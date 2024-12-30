package web

import (
	"github.com/StarJoice/tech_blog/internal/comment/internal/domain"
	"github.com/StarJoice/tech_blog/internal/comment/internal/service"
	"github.com/StarJoice/tools/ginx/session"
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
	server.POST("/comment", ginx.WithSessionAndRequest[createCommentReq](h.CreateComment))
}

func (h *Handler) PublicRoutes(server *gin.Engine) {
	//TODO implement me
	panic("implement me")
}

func (h *Handler) CreateComment(ctx *ginx.Context, req createCommentReq, sess session.Session) (ginx.Result, error) {
	uid := sess.Claims().Uid
	err := h.svc.CreateComment(ctx.Request.Context(), domain.Comment{
		Uid:           uid,
		Biz:           req.Biz,
		BizId:         req.BizId,
		Content:       req.Content,
		RootComment:   req.RootComment,
		ParentComment: req.ParentComment,
	})
	if err != nil {
		return systemErrorResult, err
	}
	return ginx.Result{
		Msg: "评论成功",
	}, err
}
