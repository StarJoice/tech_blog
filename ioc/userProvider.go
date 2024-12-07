//@Date 2024/12/5 01:07
//@Desc

package ioc

import (
	"github.com/StarJoice/tech_blog/internal/user/repository"
	"github.com/StarJoice/tech_blog/internal/user/repository/dao"
	"github.com/StarJoice/tech_blog/internal/user/service"
	"github.com/StarJoice/tech_blog/internal/user/web"
	"github.com/google/wire"
)

var UserProviderSet = wire.NewSet(
	dao.NewUserGormDao,
	repository.NewUserCacheRepository,
	service.NewUserSvc,
	web.NewUserHandle,
)
