package service

import (
	"context"
	"github.com/StarJoice/tech_blog/internal/interactive/repository"
)

type Service interface {
	LikeToggle(ctx context.Context, uid int64, biz string, id int64) error
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
