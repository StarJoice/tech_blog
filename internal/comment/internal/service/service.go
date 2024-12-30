package service

import (
	"context"
	"github.com/StarJoice/tech_blog/internal/comment/internal/domain"
	"github.com/StarJoice/tech_blog/internal/comment/internal/repository"
)

type Service interface {
	CreateComment(ctx context.Context, comment domain.Comment) error
}
type CommentService struct {
	repo repository.Repository
}

func NewCommentService(repo repository.Repository) Service {
	return &CommentService{repo: repo}
}
func (svc *CommentService) CreateComment(ctx context.Context, comment domain.Comment) error {
	return svc.repo.CreateComment(ctx, comment)
}
