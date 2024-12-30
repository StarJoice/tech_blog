package event

import (
	"github.com/StarJoice/tech_blog/internal/pkg/mqx"
	"github.com/ecodeclub/mq-api"
)

type InteractiveEvent struct {
	Biz   string `json:"biz,omitempty"`
	BizId int64  `json:"bizId,omitempty"`
	// 取值是
	// like, collect, view 三个
	Action string `json:"action,omitempty"`
	Uid    int64  `json:"uid,omitempty"`
}

type InteractiveEventProducer mqx.Producer[InteractiveEvent]

const intrTopic = "interactive_events"

func NewInteractiveEventProducer(p mq.MQ) (InteractiveEventProducer, error) {
	return mqx.NewGeneralProducer[InteractiveEvent](p, intrTopic)
}

func NewViewCntEvent(id int64, biz string) InteractiveEvent {
	return InteractiveEvent{
		Biz:    biz,
		BizId:  id,
		Action: "view",
	}
}
