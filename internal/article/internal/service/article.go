package service

import (
	"context"
	"github.com/StarJoice/tech_blog/internal/article/internal/domain"
	"github.com/StarJoice/tech_blog/internal/article/internal/event"
	"github.com/StarJoice/tech_blog/internal/article/internal/repository"
	"github.com/gotomicro/ego/core/elog"
	"golang.org/x/sync/errgroup"
	"time"
)

type Service interface {
	Save(ctx context.Context, art *domain.Article) (int64, error)
	List(ctx context.Context, offset int, limit int, uid int64) ([]domain.Article, int64, error)
	Publish(ctx context.Context, art *domain.Article) (int64, error)
	PubList(ctx context.Context, offset int, limit int) ([]domain.Article, int64, error)
	Detail(ctx context.Context, aid int64) (domain.Article, error)
	PubDetail(ctx context.Context, aid int64) (domain.Article, error)
	DelPubDetail(ctx context.Context, aid int64) error
}
type ArticleSvc struct {
	repo         repository.ArticleRepository
	intrProducer event.InteractiveEventProducer
	producer     event.SyncEventProducer
	logger       *elog.Component
	// 同步超时时长
	syncTimeout time.Duration
}

func NewArticleSvc(intrProducer event.InteractiveEventProducer,
	repo repository.ArticleRepository, producer event.SyncEventProducer) Service {
	return &ArticleSvc{
		repo:         repo,
		intrProducer: intrProducer,
		producer:     producer,
		logger:       elog.DefaultLogger,
		syncTimeout:  time.Second * 10,
	}
}

func (svc *ArticleSvc) DelPubDetail(ctx context.Context, aid int64) error {
	return svc.repo.DeletePubById(ctx, aid)
}

func (svc *ArticleSvc) PubDetail(ctx context.Context, aid int64) (domain.Article, error) {
	res, err := svc.repo.GetPubById(ctx, aid)
	if err == nil {
		go func() {
			newCtx, cancel := context.WithTimeout(context.Background(), time.Second*3)
			defer cancel()
			er := svc.intrProducer.Produce(newCtx, event.NewViewCntEvent(aid, domain.BizArticle))
			if er != nil {
				svc.logger.Error("发送阅读计数消息到消息队列失败",
					elog.FieldErr(er),
					elog.Int64("aid", aid))
			}
		}()
	}
	return res, err
}

func (svc *ArticleSvc) Detail(ctx context.Context, aid int64) (domain.Article, error) {
	return svc.repo.GetById(ctx, aid)
}

func (svc *ArticleSvc) PubList(ctx context.Context, offset int, limit int) ([]domain.Article, int64, error) {
	var (
		eg      errgroup.Group
		artList []domain.Article
		total   int64
	)
	eg.Go(func() error {
		var err error
		artList, err = svc.repo.PubList(ctx, offset, limit)
		return err
	})
	eg.Go(func() error {
		var err error
		total, err = svc.repo.PubTotal(ctx)
		return err
	})
	err := eg.Wait()
	return artList, total, err
}

func (svc *ArticleSvc) Publish(ctx context.Context, art *domain.Article) (int64, error) {
	id, err := svc.repo.Sync(ctx, art)
	if err == nil {
		go func() {
			// 同步文章数据到 搜索模块
			svc.syncArticle(id)
		}()
	}
	return id, err
}

// Save 在这里是一个upsert的语义，直接在这里分发
// 传递文章的数据可能很大，不采用值传递，使用指针（通常新建或者更新文章，数据都不为空）
func (svc *ArticleSvc) Save(ctx context.Context, art *domain.Article) (int64, error) {
	if art.Id > 0 {
		return art.Id, svc.repo.Update(ctx, art)
	}
	return svc.repo.Create(ctx, art)
}
func (svc *ArticleSvc) List(ctx context.Context,
	offset int, limit int, uid int64) ([]domain.Article, int64, error) {
	// 这里采用并发同时查询数据
	// 减少请求的整体耗时，提高性能
	var (
		eg      errgroup.Group
		artList []domain.Article
		total   int64
	)
	eg.Go(func() error {
		var err error
		artList, err = svc.repo.List(ctx, offset, limit, uid)
		return err
	})
	eg.Go(func() error {
		var err error
		total, err = svc.repo.Total(ctx, uid)
		return err
	})
	if err := eg.Wait(); err != nil {
		return artList, total, err
	}
	return artList, total, nil
	// 后续可以在这里引入缓存
}

func (svc *ArticleSvc) syncArticle(id int64) {
	ctx, cancel := context.WithTimeout(context.Background(), svc.syncTimeout)
	defer cancel()
	arts, err := svc.repo.GetPubById(ctx, id)
	if err != nil {
		svc.logger.Error("搜索案例详情失败",
			elog.FieldErr(err),
		)
		return
	}
	res := event.NewArticleEvent(arts)
	err = svc.producer.Produce(ctx, res)
	if err != nil {
		svc.logger.Error("发送案例内容到搜索失败",
			elog.FieldErr(err),
			elog.Any("article", res),
		)
	}
}
