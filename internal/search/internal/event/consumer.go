package event

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/StarJoice/tech_blog/internal/search/internal/service"
	"github.com/ecodeclub/mq-api"
	"github.com/gotomicro/ego/core/elog"
	"strconv"
	"strings"
)

type SyncConsumer struct {
	svc      service.SyncService
	consumer mq.Consumer
	logger   *elog.Component
}

func NewSyncConsumer(svc service.SyncService, q mq.MQ) (*SyncConsumer, error) {
	groupId := "sync"
	consumer, err := q.Consumer(SyncTopic, groupId)
	if err != nil {
		return nil, err
	}
	return &SyncConsumer{
		svc:      svc,
		consumer: consumer,
		logger:   elog.DefaultLogger,
	}, nil
}

func (s *SyncConsumer) Consume(ctx context.Context) error {
	msg, err := s.consumer.Consume(ctx)
	if err != nil {
		return fmt.Errorf("获取消息失败: %w", err)
	}

	var evt SyncEvent
	err = json.Unmarshal(msg.Value, &evt)
	if err != nil {
		return fmt.Errorf("解析消息失败: %w", err)
	}
	indexName := getIndexName(evt.Biz)
	docId := strconv.Itoa(evt.BizId)
	err = s.svc.Input(ctx, indexName, docId, evt.Data)
	if err != nil {
		s.logger.Error("同步消息失败", elog.Any("SyncEvent", evt))
	}
	return err
}

func (s *SyncConsumer) Start(ctx context.Context) {
	go func() {
		for {
			err := s.Consume(ctx)
			if err != nil {
				s.logger.Error("同步事件失败", elog.FieldErr(err))
			}
		}
	}()
}
func (s *SyncConsumer) Stop(_ context.Context) error {
	return s.consumer.Close()
}

func getIndexName(biz string) string {
	return fmt.Sprintf("%s_index", strings.ToLower(biz))
}
