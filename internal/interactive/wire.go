//go:build wireinject

package interactive

import (
	"context"
	"github.com/StarJoice/tech_blog/internal/interactive/internal/event"
	"github.com/StarJoice/tech_blog/internal/interactive/internal/repository"
	"github.com/StarJoice/tech_blog/internal/interactive/internal/repository/dao"
	"github.com/StarJoice/tech_blog/internal/interactive/internal/service"
	"github.com/StarJoice/tech_blog/internal/interactive/internal/web"
	"github.com/ecodeclub/mq-api"
	"github.com/ego-component/egorm"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	web.NewHandler,
	service.NewInteractiveService,
	repository.NewInteractiveRepository,
	InitDao,
)

func InitDao(db *egorm.Component) dao.InteractiveDao {
	err := dao.InitTable(db)
	if err != nil {
		panic(err)
	}
	return dao.NewInteractiveGormDao(db)
}

func InitModule(db *egorm.Component, q mq.MQ) (*Module, error) {
	wire.Build(
		ProviderSet,
		initConsumer,
		wire.Struct(new(Module), "*"),
	)
	return new(Module), nil
}

func initConsumer(svc service.Service, q mq.MQ) *event.Consumer {
	consumer, err := event.NewSyncConsumer(svc, q)
	if err != nil {
		panic(err)
	}
	consumer.Start(context.Background())
	return consumer
}
