//go:build wireinject

package search

import (
	"context"
	"github.com/StarJoice/tech_blog/internal/search/internal/event"
	"github.com/StarJoice/tech_blog/internal/search/internal/repository"
	"github.com/StarJoice/tech_blog/internal/search/internal/repository/dao"
	"github.com/StarJoice/tech_blog/internal/search/internal/service"
	"github.com/StarJoice/tech_blog/internal/search/internal/web"
	"github.com/ecodeclub/mq-api"
	"github.com/google/wire"
	"github.com/olivere/elastic/v7"
	"sync"
)

func InitModule(es *elastic.Client, q mq.MQ) (*Module, error) {
	wire.Build(
		// 初始化搜索服务
		initSearchSvc,
		// 初始化搜索的数据同步
		initSyncSvc,
		// 初始化 同步数据的消费者
		initSyncConsumer,
		web.NewHandler,
		wire.Struct(new(Module), "*"),
	)
	return new(Module), nil
}

func initSearchSvc(es *elastic.Client) service.SearchService {
	artRepo := InitRepo(es)
	return service.NewService(artRepo)
}

func InitRepo(es *elastic.Client) repository.ArticleRepo {
	InitIndexOnce(es)
	artDao := dao.NewArticleEsDao(es)
	artRepo := repository.NewArticleRepository(artDao)
	return artRepo
}

var daoOnce = sync.Once{}

func InitIndexOnce(es *elastic.Client) {
	daoOnce.Do(func() {
		err := dao.InitEs(es)
		if err != nil {
			panic(err)
		}
	})
}

func InitAnyRepo(es *elastic.Client) repository.AnyRepo {
	InitIndexOnce(es)
	anyDAO := dao.NewAnyDao(es)
	anyRepo := repository.NewAnyRepo(anyDAO)
	return anyRepo
}

func initSyncSvc(es *elastic.Client) service.SyncService {
	anyRepo := InitAnyRepo(es)
	return service.NewSyncService(anyRepo)
}
func initSyncConsumer(svc service.SyncService, q mq.MQ) *event.SyncConsumer {
	c, err := event.NewSyncConsumer(svc, q)
	if err != nil {
		panic(err)
	}
	c.Start(context.Background())
	return c
}
