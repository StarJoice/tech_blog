package web

import (
	"github.com/StarJoice/tech_blog/internal/article/domain"
	"github.com/StarJoice/tech_blog/internal/article/service"
	"github.com/StarJoice/tech_blog/pkg/slice"
	"github.com/StarJoice/tools/ginx/session"
	ginx "github.com/StarJoice/tools/ginx/wrapper"
	"github.com/gin-gonic/gin"
	"github.com/gotomicro/ego/core/elog"
)

type ArticleHandler struct {
	svc    service.Service
	logger *elog.Component
}

func NewArticleHandler(svc service.Service) *ArticleHandler {
	return &ArticleHandler{logger: elog.DefaultLogger, svc: svc}
}
func (h *ArticleHandler) PublicRoutes(server *gin.Engine) {
	// 公开的路由这里预计是未登录用户可以预览列表内容，观看详情就得登录了
	server.Group("/")
}
func (h *ArticleHandler) PrivateRoutes(server *gin.Engine) {
	art := server.Group("/article")
	art.POST("/save", ginx.WithSessionAndRequest[SaveReq](h.Save))
	// 创作中心，查看某个作者自己的帖子（无论是已经发表到线上表或者是仅仅作为草稿保存在制作表的）
	art.POST("/list", ginx.WithSessionAndRequest[Page](h.List))
}

func (h *ArticleHandler) Save(ctx *ginx.Context,
	req SaveReq, sess session.Session) (ginx.Result, error) {
	art := req.toDomain()
	art.Uid = sess.Claims().Uid
	id, err := h.svc.Save(ctx, &art)
	if err != nil {
		return systemErrorResult, err
	}
	return ginx.Result{
		Data: id,
	}, nil
}

// Page 分页
type Page struct {
	Offset int `json:"offset,omitempty" binding:"min=0"`
	Limit  int `json:"limit,omitempty" binding:"min=10"`
}

func (h *ArticleHandler) List(ctx *ginx.Context,
	req Page, sess session.Session) (ginx.Result, error) {
	// 根据uid查询
	uid := sess.Claims().Uid
	artList, total, err := h.svc.List(ctx, req.Offset, req.Limit, uid)
	if err != nil {
		return systemErrorResult, err
	}
	return ginx.Result{
		Data: h.toArtList(artList, total),
	}, nil
}

func (h *ArticleHandler) toArtList(data []domain.Article, cnt int64) ArtsList {
	return ArtsList{
		Total: cnt,
		Arts: slice.Map(data, func(idx int, art domain.Article) Article {
			return newArt(art)
		}),
	}
}

func newArt(art domain.Article) Article {
	return Article{
		Id:      art.Id,
		Title:   art.Title,
		Content: art.Content,
		Ctime:   art.Ctime,
	}
}
