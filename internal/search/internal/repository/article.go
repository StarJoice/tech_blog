package repository

import (
	"context"
	"github.com/StarJoice/tech_blog/internal/search/internal/domain"
	"github.com/StarJoice/tech_blog/internal/search/internal/repository/dao"
	"time"
)

type ArticleRepo interface {
	SearchArticle(tx context.Context, offset, limit int, queryMetas []domain.QueryMeta) ([]domain.Article, error)
}
type ArticleRepository struct {
	ArtDao dao.ArticleDao
}

func NewArticleRepository(artDao dao.ArticleDao) ArticleRepo {
	return &ArticleRepository{ArtDao: artDao}
}
func (repo *ArticleRepository) SearchArticle(tx context.Context, offset, limit int, queryMetas []domain.QueryMeta) ([]domain.Article, error) {
	arts, err := repo.ArtDao.SearchArticle(tx, offset, limit, queryMetas)
	if err != nil {
		return nil, err
	}
	Articles := make([]domain.Article, 0, len(arts))
	for _, art := range arts {
		Articles = append(Articles, repo.toDomain(art))
	}
	return Articles, err
}

func (repo *ArticleRepository) toDomain(a dao.Article) domain.Article {
	return domain.Article{
		Id:       a.Id,
		Uid:      a.Uid,
		Title:    a.Title,
		Content:  a.Content,
		Abstract: a.Abstract,
		Ctime:    time.UnixMilli(a.Ctime),
		Utime:    time.UnixMilli(a.Utime),
	}
}
