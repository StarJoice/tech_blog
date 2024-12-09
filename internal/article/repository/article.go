//@Date 2024/12/9 17:03
//@Desc

package repository

import (
	"context"
	"github.com/StarJoice/tech_blog/internal/article/domain"
	"github.com/StarJoice/tech_blog/internal/article/repository/dao"
)

type ArticleRepository interface {
	// 创作者接口
	Create(ctx context.Context, art *domain.Article) (int64, error)
	Update(ctx context.Context, art *domain.Article) error
	List(ctx context.Context, offset int, limit int, uid int64) ([]domain.Article, error)
	Total(ctx context.Context, uid int64) (int64, error)
}

type ArticleCachedRepository struct {
	dao dao.ArticleDao
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
		Ctime:   art.Ctime,
	}
}
