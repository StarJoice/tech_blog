package repository

import (
	"context"
	"github.com/StarJoice/tech_blog/internal/search/internal/repository/dao"
)

type AnyRepo interface {
	Input(ctx context.Context, index string, docID string, data string) error
}

type anyRepo struct {
	anyDao dao.AnyDao
}

func NewAnyRepo(anyDao dao.AnyDao) AnyRepo {
	return &anyRepo{
		anyDao: anyDao,
	}
}
func (a *anyRepo) Input(ctx context.Context, index string, docID string, data string) error {
	return a.anyDao.Input(ctx, index, docID, data)
}
