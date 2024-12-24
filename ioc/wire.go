//go:build wireinject

package ioc

import (
	"github.com/StarJoice/tech_blog/internal/article"
	"github.com/StarJoice/tech_blog/internal/interactive"
	"github.com/StarJoice/tech_blog/internal/label"
	"github.com/StarJoice/tech_blog/internal/user"
	"github.com/google/wire"
)

var BaseSet = wire.NewSet(InitDB, InitSession, InitRedis, InitMq)

func InitApp() (*App, error) {
	wire.Build(wire.Struct(new(App), "*"),
		BaseSet,
		user.InitModule,
		wire.FieldsOf(new(*user.Module), "Hdl"),
		article.InitModule,
		wire.FieldsOf(new(*article.Module), "Hdl"),
		label.InitModule,
		wire.FieldsOf(new(*label.Module), "Hdl"),
		interactive.InitModule,
		wire.FieldsOf(new(*interactive.Module), "Hdl"),
		InitGinXServer)
	return new(App), nil
}
