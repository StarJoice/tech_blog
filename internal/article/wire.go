//go:build wireinject

package article

import (
	"github.com/StarJoice/tech_blog/internal/article/repository"
	"github.com/StarJoice/tech_blog/internal/article/repository/dao"
	"github.com/StarJoice/tech_blog/internal/article/service"
	"github.com/StarJoice/tech_blog/internal/article/web"
	"github.com/ego-component/egorm"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	web.NewArticleHandler,
	repository.NewArticleCachedRepository,
	service.NewArticleSvc,
	InitDao,
)

func InitDao(db *egorm.Component) dao.ArticleDao {
	err := dao.InitTable(db)
	if err != nil {
		panic(err)
	}
	return dao.NewArticleGormDao(db)
}

func InitHandler(db *egorm.Component) *Handler {
	wire.Build(ProviderSet)
	return new(Handler)
}

type Handler = web.ArticleHandler
