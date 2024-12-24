package service

import (
	"context"
	"github.com/StarJoice/tech_blog/internal/interactive/internal/repository"
)

type Service interface {
	LikeToggle(ctx context.Context, uid int64, biz string, id int64) error
	IncrReadCnt(ctx context.Context, biz string, bizId int64) error
	CollectToggle(ctx context.Context, biz string, id int64, uid int64) error
}

type InteractiveService struct {
	repo repository.Repository
}

func NewInteractiveService(repo repository.Repository) Service {
	return &InteractiveService{repo: repo}
}

func (svc *InteractiveService) LikeToggle(ctx context.Context, uid int64, biz string, id int64) error {
	return svc.repo.LikeToggle(ctx, uid, biz, id)
}

func (svc *InteractiveService) IncrReadCnt(ctx context.Context, biz string, bizId int64) error {
	return svc.repo.IncrViewCnt(ctx, biz, bizId)
}
func (svc *InteractiveService) CollectToggle(ctx context.Context, biz string, id int64, uid int64) error {
	return svc.repo.CollectToggle(ctx, biz, id, uid)
}
