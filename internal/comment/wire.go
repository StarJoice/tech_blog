//go:build wireinject

package comment

import (
	"github.com/StarJoice/tech_blog/internal/comment/internal/repository"
	"github.com/StarJoice/tech_blog/internal/comment/internal/repository/dao"
	"github.com/StarJoice/tech_blog/internal/comment/internal/service"
	"github.com/StarJoice/tech_blog/internal/comment/internal/web"
	"github.com/ego-component/egorm"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	web.NewHandler,
	service.NewCommentService,
	repository.NewCommentRepository,
	InitDao,
)

func InitDao(db *egorm.Component) dao.CommentDao {
	err := dao.InitTable(db)
	if err != nil {
		panic(err)
	}
	return dao.NewCommentDao(db)
}
func InitModule(db *egorm.Component) (*Module, error) {
	wire.Build(
		ProviderSet,
		wire.Struct(new(Module), "*"),
	)
	return new(Module), nil
}
