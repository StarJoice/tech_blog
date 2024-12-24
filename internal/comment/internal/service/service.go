package service

import "github.com/StarJoice/tech_blog/internal/comment/internal/repository"

type Service interface {
}
type CommentService struct {
	repo repository.Repository
}
