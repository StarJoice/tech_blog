package event

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/StarJoice/tech_blog/internal/interactive/internal/service"
	"github.com/ecodeclub/mq-api"
	"github.com/gotomicro/ego/core/elog"
)

const topic = "interactive_events"

type Consumer struct {
	handlerMap map[string]handleFunc
	consumer   mq.Consumer
	svc        service.Service
	logger     *elog.Component
}

func NewSyncConsumer(svc service.Service, q mq.MQ) (*Consumer, error) {
	groupID := "interactive_group"
	consumer, err := q.Consumer(topic, groupID)
	if err != nil {
		return nil, err
	}
	c := &Consumer{
		consumer: consumer,
		svc:      svc,
		logger:   elog.DefaultLogger,
	}
	handlerMap := map[string]handleFunc{
		"like":    c.likeHandle,
		"collect": c.collectHandle,
		"view":    c.viewHandle,
	}
	c.handlerMap = handlerMap
	return c, nil
}

func (c *Consumer) likeHandle(ctx context.Context, svc service.Service, evt Event) error {
	return svc.LikeToggle(ctx, evt.Uid, evt.Biz, evt.BizId)
}
func (c *Consumer) collectHandle(ctx context.Context, svc service.Service, evt Event) error {
	return svc.CollectToggle(ctx, evt.Biz, evt.BizId, evt.Uid)
}
func (c *Consumer) viewHandle(ctx context.Context, svc service.Service, evt Event) error {
	return svc.IncrReadCnt(ctx, evt.Biz, evt.BizId)

}
func (c *Consumer) Consume(ctx context.Context) error {
	msg, err := c.consumer.Consume(ctx)
	if err != nil {
		return fmt.Errorf("获取消息失败: %w", err)
	}

	var evt Event
	err = json.Unmarshal(msg.Value, &evt)
	if err != nil {
		return fmt.Errorf("解析消息失败: %w", err)
	}
	handler, ok := c.handlerMap[evt.Action]
	if !ok {
		return errors.New("未找到相关业务的处理方法")
	}
	err = handler(ctx, c.svc, evt)
	if err != nil {
		c.logger.Error("同步消息失败", elog.Any("interactive_event", evt))
	}
	return err
}

func (c *Consumer) Start(ctx context.Context) {
	go func() {
		for {
			err := c.Consume(ctx)
			if err != nil {
				c.logger.Error("同步事件失败", elog.FieldErr(err))
			}
		}
	}()
}
func (c *Consumer) Stop(_ context.Context) error {
	return c.consumer.Close()
}
