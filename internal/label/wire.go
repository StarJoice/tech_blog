//go:build wireinject

package label

import (
	"github.com/StarJoice/tech_blog/internal/label/internal/repository"
	"github.com/StarJoice/tech_blog/internal/label/internal/repository/dao"
	"github.com/StarJoice/tech_blog/internal/label/internal/service"
	"github.com/StarJoice/tech_blog/internal/label/internal/web"
	"github.com/ego-component/egorm"
	"github.com/google/wire"
)

func InitDao(db *egorm.Component) dao.LabelDao {
	err := dao.InitTable(db)
	if err != nil {
		panic(err)
	}
	return dao.NewLabelGormDao(db)
}

var ProviderSet = wire.NewSet(
	web.NewHandler,
	service.NewLabelService,
	repository.NewLabelRepository,
	InitDao,
)

func InitModule(db *egorm.Component) (*Module, error) {
	wire.Build(
		ProviderSet,
		wire.Struct(new(Module), "*"),
	)
	return new(Module), nil
}
