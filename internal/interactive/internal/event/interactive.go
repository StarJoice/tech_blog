package event

import (
	"context"
	"github.com/StarJoice/tech_blog/internal/interactive/internal/service"
)

type Event struct {
	Biz   string `json:"biz,omitempty"`
	BizId int64  `json:"bizId,omitempty"`
	// 取值是
	// like, collect, view 三个
	Action string `json:"action,omitempty"`
	Uid    int64  `json:"uid,omitempty"`
}

type handleFunc func(ctx context.Context, svc service.Service, evt Event) error
