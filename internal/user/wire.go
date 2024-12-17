//go:build wireinject

package user

import (
	"github.com/StarJoice/tech_blog/internal/user/internal/repository"
	"github.com/StarJoice/tech_blog/internal/user/internal/repository/cache"
	"github.com/StarJoice/tech_blog/internal/user/internal/repository/dao"
	"github.com/StarJoice/tech_blog/internal/user/internal/service"
	"github.com/StarJoice/tech_blog/internal/user/internal/web"
	"github.com/ego-component/egorm"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
)

var ProviderSet = wire.NewSet(
	cache.NewUserRedisCache,
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

func InitModule(db *egorm.Component, cmd redis.Cmdable) (*Module, error) {
	wire.Build(
		ProviderSet,
		wire.Struct(new(Module), "*"),
	)
	return new(Module), nil
}
