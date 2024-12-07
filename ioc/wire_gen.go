// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package ioc

import (
	"github.com/StarJoice/tech_blog/internal/user"
	"github.com/google/wire"
)

// Injectors from wire.go:

func InitApp() (*App, error) {
	cmdable := InitRedis()
	provider := InitSession(cmdable)
	db := InitDB()
	userHandler := user.InitHandler(db)
	component := InitGinXServer(provider, userHandler)
	app := &App{
		Web: component,
	}
	return app, nil
}

// wire.go:

var BaseSet = wire.NewSet(InitDB, InitSession, InitRedis)
