package repository

import (
	"context"
	"github.com/StarJoice/tech_blog/internal/interactive/repository/dao"
)

type Repository interface {
	LikeToggle(ctx context.Context, uid int64, biz string, id int64) error
}

type InteractiveRepository struct {
	dao dao.InteractiveDao
}

func NewInteractiveRepository(dao dao.InteractiveDao) Repository {
	return &InteractiveRepository{dao: dao}
}

func (repo *InteractiveRepository) LikeToggle(ctx context.Context, uid int64, biz string, id int64) error {
	return repo.dao.LikeToggle(ctx, uid, biz, id)
}
