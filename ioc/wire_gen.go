// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package ioc

import (
	"github.com/StarJoice/tech_blog/internal/article"
	"github.com/StarJoice/tech_blog/internal/user"
	"github.com/google/wire"
)

import (
	_ "github.com/StarJoice/tech_blog/docs"
)

// Injectors from wire.go:

func InitApp() (*App, error) {
	cmdable := InitRedis()
	provider := InitSession(cmdable)
	db := InitDB()
	module, err := user.InitModule(db, cmdable)
	if err != nil {
		return nil, err
	}
	userHandler := module.Hdl
	articleModule, err := article.InitModule(db, module)
	if err != nil {
		return nil, err
	}
	articleHandler := articleModule.Hdl
	component := InitGinXServer(provider, userHandler, articleHandler)
	app := &App{
		Web: component,
	}
	return app, nil
}

// wire.go:

var BaseSet = wire.NewSet(InitDB, InitSession, InitRedis)
