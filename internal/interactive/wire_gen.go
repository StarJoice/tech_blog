// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package interactive

import (
	"github.com/StarJoice/tech_blog/internal/interactive/repository"
	"github.com/StarJoice/tech_blog/internal/interactive/repository/dao"
	"github.com/StarJoice/tech_blog/internal/interactive/service"
	"github.com/StarJoice/tech_blog/internal/interactive/web"
	"github.com/ego-component/egorm"
	"github.com/google/wire"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func InitModule(db *gorm.DB) (*Module, error) {
	interactiveDao := InitDao(db)
	repositoryRepository := repository.NewInteractiveRepository(interactiveDao)
	serviceService := service.NewInteractiveService(repositoryRepository)
	handler := web.NewHandler(serviceService)
	module := &Module{
		Hdl: handler,
		Svc: serviceService,
	}
	return module, nil
}

// wire.go:

var ProviderSet = wire.NewSet(web.NewHandler, service.NewInteractiveService, repository.NewInteractiveRepository, InitDao)

func InitDao(db *egorm.Component) dao.InteractiveDao {
	err := dao.InitTable(db)
	if err != nil {
		panic(err)
	}
	return dao.NewInteractiveGormDao(db)
}
