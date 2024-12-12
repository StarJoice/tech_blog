//go:build wireinject

package article

import (
	"github.com/StarJoice/tech_blog/internal/article/repository"
	"github.com/StarJoice/tech_blog/internal/article/repository/dao"
	"github.com/StarJoice/tech_blog/internal/article/service"
	"github.com/StarJoice/tech_blog/internal/article/web"
	"github.com/StarJoice/tech_blog/internal/user"
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

func InitModule(db *egorm.Component, u *user.Module) (*Module, error) {
	wire.Build(
		ProviderSet,
		wire.FieldsOf(new(*user.Module), "Svc"),
		wire.Struct(new(Module), "*"),
	)
	return new(Module), nil
}
