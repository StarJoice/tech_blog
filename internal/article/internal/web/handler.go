package web

import (
	"github.com/StarJoice/tech_blog/internal/article/internal/domain"
	"github.com/StarJoice/tech_blog/internal/article/internal/service"
	"github.com/StarJoice/tech_blog/internal/user"
	"github.com/StarJoice/tech_blog/pkg/slice"
	"github.com/StarJoice/tools/ginx/session"
	ginx "github.com/StarJoice/tools/ginx/wrapper"
	"github.com/gin-gonic/gin"
	"github.com/gotomicro/ego/core/elog"
	"time"
)

type ArticleHandler struct {
	svc    service.Service
	logger *elog.Component
	// 通过组合userService 来实现查询作者？
	UserSvc user.Service
}

func NewArticleHandler(svc service.Service, UserSvc user.Service) *ArticleHandler {
	return &ArticleHandler{logger: elog.DefaultLogger, svc: svc, UserSvc: UserSvc}
}
func (h *ArticleHandler) PublicRoutes(server *gin.Engine) {
	// 公开的路由这里预计是未登录用户可以预览列表内容，观看详情就得登录了
	art := server.Group("/article")
	art.POST("/publish/list", ginx.WithRequest[Page](h.PubList))
}
func (h *ArticleHandler) PrivateRoutes(server *gin.Engine) {
	art := server.Group("/article")
	art.POST("/save", ginx.WithSessionAndRequest[SaveReq](h.Save))
	// 创作中心，查看某个作者自己的帖子（无论是已经发表或者是仅仅作为草稿保存在制作表的）
	art.POST("/list", ginx.WithSessionAndRequest[Page](h.List))
	art.POST("/publish", ginx.WithSessionAndRequest[SaveReq](h.Publish))
	art.POST("/detail", ginx.WithRequest[ArtId](h.Detail))
	art.POST("/publish/detail", ginx.WithRequest[ArtId](h.PubDetail))
	// 删除自己已发表的帖子
	// todo 暂时先这样实现，后续可能软删除-- 即实现线上文章下线，回到创作中心继续编辑重新上线的功能
	art.DELETE("/publish/detail", ginx.WithRequest[ArtId](h.DelPubDetail))
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

func (h *ArticleHandler) Publish(ctx *ginx.Context,
	req SaveReq, sess session.Session) (ginx.Result, error) {
	art := req.toDomain()
	art.Uid = sess.Claims().Uid
	id, err := h.svc.Publish(ctx, &art)
	if err != nil {
		return systemErrorResult, err
	}
	return ginx.Result{
		Data: id,
	}, nil
}

func (h *ArticleHandler) PubList(ctx *ginx.Context,
	req Page) (ginx.Result, error) {
	data, cnt, err := h.svc.PubList(ctx, req.Offset, req.Limit)
	if err != nil {
		return systemErrorResult, err
	}
	authors := make(map[int64]Author)
	// 遍历每篇文章，查询作者信息
	for _, d := range data {
		if _, exists := authors[d.Uid]; !exists {
			u, err := h.UserSvc.GetByID(ctx, d.Uid)
			if err != nil {
				return systemErrorResult, err
			}
			authors[d.Uid] = Author{
				Nickname: u.Nickname,
				Avatar:   u.Avatar,
			}
		}
	}
	return ginx.Result{
		Data: ArtsList{
			Arts: slice.Map(data, func(idx int, art domain.Article) Article {
				return Article{
					Id:     art.Id,
					Uid:    art.Uid,
					Title:  art.Title,
					Ctime:  art.Utime.Format(time.DateTime),
					Author: authors[art.Uid],
				}
			}),
			Total: cnt,
		},
	}, nil
}

func (h *ArticleHandler) Detail(ctx *ginx.Context,
	req ArtId) (ginx.Result, error) {
	detail, err := h.svc.Detail(ctx, req.Aid)
	if err != nil {
		return systemErrorResult, err
	}
	return ginx.Result{
		Data: newArt(detail),
	}, err
}

func (h *ArticleHandler) PubDetail(ctx *ginx.Context,
	req ArtId) (ginx.Result, error) {
	detail, err := h.svc.PubDetail(ctx, req.Aid)
	if err != nil {
		return systemErrorResult, err
	}
	return ginx.Result{
		Data: newArt(detail),
	}, err
}

func (h *ArticleHandler) DelPubDetail(ctx *ginx.Context, req ArtId) (ginx.Result, error) {
	err := h.svc.DelPubDetail(ctx, req.Aid)
	if err != nil {
		return systemErrorResult, err
	}
	return ginx.Result{
		Msg: "删除成功",
	}, err
}

func newArt(art domain.Article) Article {
	return Article{
		Id:       art.Id,
		Title:    art.Title,
		Abstract: art.Abstract,
		Ctime:    art.Ctime.Format(time.DateTime),
	}
}
