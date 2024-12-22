package service

import "github.com/StarJoice/tech_blog/internal/label/repository"

type Service interface {
}
type LabelService struct {
	repo repository.Repository
}

func NewLabelService(repo repository.Repository) Service {
	return &LabelService{repo: repo}
}
