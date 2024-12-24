//go:build wireinject

package article

import (
	"github.com/StarJoice/tech_blog/internal/article/internal/event"
	"github.com/StarJoice/tech_blog/internal/article/internal/repository"
	"github.com/StarJoice/tech_blog/internal/article/internal/repository/dao"
	"github.com/StarJoice/tech_blog/internal/article/internal/service"
	"github.com/StarJoice/tech_blog/internal/article/internal/web"
	"github.com/StarJoice/tech_blog/internal/user"
	"github.com/ecodeclub/mq-api"
	"github.com/ego-component/egorm"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	web.NewArticleHandler,
	repository.NewArticleCachedRepository,
	service.NewArticleSvc,
	event.NewInteractiveEventProducer,
	InitDao,
)

func InitDao(db *egorm.Component) dao.ArticleDao {
	err := dao.InitTable(db)
	if err != nil {
		panic(err)
	}
	return dao.NewArticleGormDao(db)
}

func InitModule(db *egorm.Component, u *user.Module, q mq.MQ) (*Module, error) {
	wire.Build(
		ProviderSet,
		wire.FieldsOf(new(*user.Module), "Svc"),
		wire.Struct(new(Module), "*"),
	)
	return new(Module), nil
}
