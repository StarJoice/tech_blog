//go:build wireinject

package startup

import (
	"github.com/StarJoice/tech_blog/internal/user"
	"github.com/StarJoice/tech_blog/internal/user/internal/repository"
	"github.com/StarJoice/tech_blog/internal/user/internal/repository/dao"
	"github.com/StarJoice/tech_blog/internal/user/internal/service"
	"github.com/StarJoice/tech_blog/internal/user/internal/web"
	"github.com/StarJoice/tech_blog/ioc"
	"github.com/google/wire"
)

func Inithandler() *user.Handler {
	wire.Build(
		ioc.InitDB,
		web.NewUserHandle,
		service.NewUserSvc,
		repository.NewUserCacheRepository,
		dao.NewUserGormDao,
	)
	return new(user.Handler)
}
