package event

import (
	"context"
	"github.com/StarJoice/tech_blog/internal/pkg/mqx"
	"github.com/ecodeclub/mq-api"
)

// SyncTopic 生产者TOPIC
const SyncTopic = "sync_data_to_search"

type SyncEventProducer interface {
	Produce(ctx context.Context, evt ArticleEvent) error
}

func NewSyncEventProducer(q mq.MQ) (SyncEventProducer, error) {
	return mqx.NewGeneralProducer[ArticleEvent](q, SyncTopic)
}
