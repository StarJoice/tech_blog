package service

import (
	"context"
	"github.com/StarJoice/tech_blog/internal/article/domain"
	"github.com/StarJoice/tech_blog/internal/article/repository"
	"golang.org/x/sync/errgroup"
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
	repo repository.ArticleRepository
}

func (svc *ArticleSvc) DelPubDetail(ctx context.Context, aid int64) error {
	return svc.repo.DeletePubById(ctx, aid)
}

func (svc *ArticleSvc) PubDetail(ctx context.Context, aid int64) (domain.Article, error) {
	return svc.repo.GetPubById(ctx, aid)
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
	return svc.repo.Sync(ctx, art)
}

func NewArticleSvc(repo repository.ArticleRepository) Service {
	return &ArticleSvc{repo: repo}
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
