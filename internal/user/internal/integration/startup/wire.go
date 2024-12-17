//go:build wireinject

package startup

import (
	testioc "github.com/StarJoice/tech_blog/internal/test/ioc"
	"github.com/StarJoice/tech_blog/internal/user"
	"github.com/StarJoice/tech_blog/internal/user/internal/repository"
	"github.com/StarJoice/tech_blog/internal/user/internal/repository/dao"
	"github.com/StarJoice/tech_blog/internal/user/internal/service"
	"github.com/StarJoice/tech_blog/internal/user/internal/web"
	"github.com/google/wire"
)

func InitHandler() *user.Handler {
	wire.Build(
		testioc.BaseSet,
		web.NewUserHandle,
		service.NewUserSvc,
		repository.NewUserCacheRepository,
		dao.NewUserGormDao,
	)
	return new(user.Handler)
}
