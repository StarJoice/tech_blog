//go:build wireinject

package user

import (
	"github.com/StarJoice/tech_blog/internal/user/internal/repository"
	"github.com/StarJoice/tech_blog/internal/user/internal/repository/dao"
	"github.com/StarJoice/tech_blog/internal/user/internal/service"
	"github.com/StarJoice/tech_blog/internal/user/internal/web"
	"github.com/ego-component/egorm"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	web.NewUserHandle,
	service.NewUserSvc,
	repository.NewUserCacheRepository,
	InitDao,
)

func InitDao(db *egorm.Component) dao.UserDao {
	err := dao.InitTable(db)
	if err != nil {
		panic(err)
	}
	return dao.NewUserGormDao(db)
}

func InitHandler(db *egorm.Component) *Handler {
	wire.Build(ProviderSet)
	return new(Handler)
}

type Handler = web.UserHandler
