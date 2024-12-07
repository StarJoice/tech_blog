//go:build wireinject

package ioc

import "github.com/google/wire"

func InitApp() (*App, error) {
	wire.Build(wire.Struct(new(App), "*"),
		InitDB, InitSession, InitRedis,
		UserProviderSet,
		InitGinXServer)
	return new(App), nil
}
