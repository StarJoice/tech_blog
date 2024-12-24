package service

import "github.com/StarJoice/tech_blog/internal/comment/internal/repository"

type Service interface {
}
type CommentService struct {
	repo repository.Repository
}

func NewCommentService(repo repository.Repository) Service {
	return &CommentService{repo: repo}
}
