//go:build wireinject

package ioc

import (
	"github.com/StarJoice/tech_blog/internal/article"
	"github.com/StarJoice/tech_blog/internal/user"
	"github.com/google/wire"
)

var BaseSet = wire.NewSet(InitDB, InitSession, InitRedis)

func InitApp() (*App, error) {
	wire.Build(wire.Struct(new(App), "*"),
		BaseSet,
		user.InitHandler,
		article.InitHandler,
		InitGinXServer)
	return new(App), nil
}
