package service

import (
	"context"
	"github.com/StarJoice/tech_blog/internal/search/internal/repository"
)

// SyncService 用来同步数据
type SyncService interface {
	Input(ctx context.Context, index string, docID string, data string) error
}

type syncService struct {
	anyRepo repository.AnyRepo
}

func NewSyncService(anyRepo repository.AnyRepo) SyncService {
	return &syncService{anyRepo: anyRepo}
}

func (s *syncService) Input(ctx context.Context, index string, docID string, data string) error {
	return s.anyRepo.Input(ctx, index, docID, data)
}
