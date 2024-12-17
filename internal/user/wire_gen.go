// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

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
	"gorm.io/gorm"
)

// Injectors from wire.go:

func InitModule(db *gorm.DB, cmd redis.Cmdable) (*Module, error) {
	userDao := InitDao(db)
	userCache := cache.NewUserRedisCache(cmd)
	userRepository := repository.NewUserCacheRepository(userDao, userCache)
	userService := service.NewUserSvc(userRepository)
	userHandler := web.NewUserHandle(userService)
	module := &Module{
		Svc: userService,
		Hdl: userHandler,
	}
	return module, nil
}

// wire.go:

var ProviderSet = wire.NewSet(cache.NewUserRedisCache, web.NewUserHandle, service.NewUserSvc, repository.NewUserCacheRepository, InitDao)

func InitDao(db *egorm.Component) dao.UserDao {
	err := dao.InitTable(db)
	if err != nil {
		panic(err)
	}
	return dao.NewUserGormDao(db)
}
