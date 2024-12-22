//go:build wireinject

package interactive

import (
	"github.com/StarJoice/tech_blog/internal/interactive/repository"
	"github.com/StarJoice/tech_blog/internal/interactive/repository/dao"
	"github.com/StarJoice/tech_blog/internal/interactive/service"
	"github.com/StarJoice/tech_blog/internal/interactive/web"
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

func InitModule(db *egorm.Component) (*Module, error) {
	wire.Build(
		ProviderSet,
		wire.Struct(new(Module), "*"),
	)
	return new(Module), nil
}
