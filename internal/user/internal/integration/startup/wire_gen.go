// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package startup

import (
	"github.com/StarJoice/tech_blog/internal/user/internal/repository"
	"github.com/StarJoice/tech_blog/internal/user/internal/repository/dao"
	"github.com/StarJoice/tech_blog/internal/user/internal/service"
	"github.com/StarJoice/tech_blog/internal/user/internal/web"
	"github.com/StarJoice/tech_blog/ioc"
)

// Injectors from wire.go:

func Inithandler() *web.UserHandler {
	db := ioc.InitDB()
	userDao := dao.NewUserGormDao(db)
	userRepository := repository.NewUserCacheRepository(userDao)
	userService := service.NewUserSvc(userRepository)
	userHandler := web.NewUserHandle(userService)
	return userHandler
}
