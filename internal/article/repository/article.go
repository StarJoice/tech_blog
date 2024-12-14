//@Date 2024/12/9 17:03
//@Desc

package repository

import (
	"context"
	"github.com/StarJoice/tech_blog/internal/article/domain"
	"github.com/StarJoice/tech_blog/internal/article/repository/dao"
	"time"
)

type ArticleRepository interface {
	// 创作者接口
	Create(ctx context.Context, art *domain.Article) (int64, error)
	Update(ctx context.Context, art *domain.Article) error
	List(ctx context.Context, offset int, limit int, uid int64) ([]domain.Article, error)
	Total(ctx context.Context, uid int64) (int64, error)
	Sync(ctx context.Context, art *domain.Article) (int64, error)
	// 线上库查询
	PubList(ctx context.Context, offset int, limit int) ([]domain.Article, error)
	PubTotal(ctx context.Context) (int64, error)
	GetById(ctx context.Context, aid int64) (domain.Article, error)
	GetPubById(ctx context.Context, aid int64) (domain.Article, error)
	DeletePubById(ctx context.Context, aid int64) error
}

type ArticleCachedRepository struct {
	dao dao.ArticleDao
}

func (repo *ArticleCachedRepository) DeletePubById(ctx context.Context, aid int64) error {
	return repo.dao.DeleteById(ctx, aid)
}

func (repo *ArticleCachedRepository) GetPubById(ctx context.Context, aid int64) (domain.Article, error) {
	data, err := repo.dao.GetPubArtById(ctx, aid)
	if err != nil {
		return domain.Article{}, err
	}
	return repo.toDomain(dao.Article(data)), err
}

func (repo *ArticleCachedRepository) GetById(ctx context.Context, aid int64) (domain.Article, error) {
	date, err := repo.dao.GetArtById(ctx, aid)
	if err != nil {
		return domain.Article{}, err
	}
	return repo.toDomain(date), err
}

func (repo *ArticleCachedRepository) PubList(ctx context.Context,
	offset int, limit int) ([]domain.Article, error) {
	artList, err := repo.dao.PubList(ctx, offset, limit)
	if err != nil {
		return nil, err
	}
	domainArts := make([]domain.Article, 0, len(artList))
	for _, art := range artList {
		domainArts = append(domainArts, repo.toDomain(dao.Article(art)))
	}
	return domainArts, nil
}

func (repo *ArticleCachedRepository) PubTotal(ctx context.Context) (int64, error) {
	res, err := repo.dao.PublishTotal(ctx)
	if err != nil {
		return 0, err
	}
	return res, nil
}

// Sync 这里的语义是同步（将文章发表到线上，即同步数据到线上表）
func (repo *ArticleCachedRepository) Sync(ctx context.Context, art *domain.Article) (int64, error) {
	data := repo.toEneity(art)
	return repo.dao.Sync(ctx, data)
}

func (repo *ArticleCachedRepository) Total(ctx context.Context,
	uid int64) (int64, error) {
	return repo.dao.Count(ctx, uid)
}

func (repo *ArticleCachedRepository) List(ctx context.Context,
	offset int, limit int, uid int64) ([]domain.Article, error) {
	artList, err := repo.dao.List(ctx, offset, limit, uid)
	if err != nil {
		return nil, err
	}
	domainArts := make([]domain.Article, 0, len(artList))
	for _, art := range artList {
		domainArts = append(domainArts, repo.toDomain(art))
	}
	return domainArts, nil
}

func (repo *ArticleCachedRepository) Create(ctx context.Context, art *domain.Article) (int64, error) {
	return repo.dao.Create(ctx, repo.toEneity(art))
}

func (repo *ArticleCachedRepository) Update(ctx context.Context, art *domain.Article) error {
	return repo.dao.Update(ctx, repo.toEneity(art))
}

func NewArticleCachedRepository(dao dao.ArticleDao) ArticleRepository {
	return &ArticleCachedRepository{dao: dao}
}

func (repo *ArticleCachedRepository) toEneity(art *domain.Article) dao.Article {
	return dao.Article{
		Id:      art.Id,
		Uid:     art.Uid,
		Title:   art.Title,
		Content: art.Content,
	}
}

func (repo *ArticleCachedRepository) toDomain(art dao.Article) domain.Article {
	return domain.Article{
		Id:      art.Id,
		Uid:     art.Uid,
		Title:   art.Title,
		Content: art.Content,
		Ctime:   time.UnixMilli(art.Ctime),
	}
}
