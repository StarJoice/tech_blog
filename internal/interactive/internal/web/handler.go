package web

import (
	"github.com/StarJoice/tech_blog/internal/interactive/internal/service"
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

// PrivateRoutes 这里暂时让前端控制biz和bizid（反范式）
func (h Handler) PrivateRoutes(server *gin.Engine) {
	server.POST("/like", ginx.WithSessionAndRequest[likeReq](h.Like))
}

func (h Handler) PublicRoutes(server *gin.Engine) {
	panic("implement me")
}

func (h Handler) Like(ctx *ginx.Context, req likeReq,
	sess session.Session) (ginx.Result, error) {
	uid := sess.Claims().Uid
	err := h.svc.LikeToggle(ctx, uid, req.Biz, req.BizId)
	if err != nil {
		return ginx.Result{}, err
	}
	return ginx.Result{
		Msg: "点赞成功",
	}, nil
}
