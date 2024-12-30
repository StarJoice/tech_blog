package service

import (
	"context"
	"github.com/StarJoice/tech_blog/internal/search/internal/domain"
	"github.com/StarJoice/tech_blog/internal/search/internal/repository"
)

type SearchHandler interface {
	search(ctx context.Context, metas []domain.QueryMeta, offset, limit int, res *domain.SearchResult) error
}

type articleHandler struct {
	repo repository.ArticleRepo
}

func NewArticleHandler(repo repository.ArticleRepo) SearchHandler {
	return &articleHandler{repo: repo}
}

func (a *articleHandler) search(ctx context.Context,
	metas []domain.QueryMeta, offset, limit int, res *domain.SearchResult) error {
	arts, err := a.repo.SearchArticle(ctx, offset, limit, metas)
	if err != nil {
		return err
	}
	res.SetArticles(arts)
	return nil
}
