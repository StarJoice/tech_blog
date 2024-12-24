package repository

import (
	"context"
	"errors"
	dao2 "github.com/StarJoice/tech_blog/internal/interactive/internal/repository/dao"
)

type Repository interface {
	LikeToggle(ctx context.Context, uid int64, biz string, id int64) error
	IncrViewCnt(ctx context.Context, biz string, bizId int64) error
	Liked(ctx context.Context, biz string, id int64, uid int64) (bool, error)
	CollectToggle(ctx context.Context, biz string, id int64, uid int64) error
}

type InteractiveRepository struct {
	dao dao2.InteractiveDao
}

func NewInteractiveRepository(dao dao2.InteractiveDao) Repository {
	return &InteractiveRepository{dao: dao}
}

func (repo *InteractiveRepository) LikeToggle(ctx context.Context, uid int64, biz string, id int64) error {
	return repo.dao.LikeToggle(ctx, uid, biz, id)
}
func (repo *InteractiveRepository) IncrViewCnt(ctx context.Context, biz string, bizId int64) error {
	return repo.dao.IncrViewCnt(ctx, biz, bizId)
}
func (repo *InteractiveRepository) Liked(ctx context.Context, biz string, id int64, uid int64) (bool, error) {
	_, err := repo.dao.GetLikeInfo(ctx, id, biz, uid)
	switch {
	case err == nil:
		return true, nil
	case errors.Is(err, dao2.ErrRecordNotFound):
		return false, nil
	default:
		return false, err
	}
}
func (repo *InteractiveRepository) CollectToggle(ctx context.Context, biz string, id int64, uid int64) error {
	return repo.dao.CollectToggle(ctx, dao2.UserCollectionBiz{
		Uid:   uid,
		BizId: id,
		Biz:   biz,
	})
}
