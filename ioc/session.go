//@Date 2024/12/7 15:43
//@Desc

package ioc

import (
	"github.com/StarJoice/tools/ginx/session"
	redis2 "github.com/StarJoice/tools/ginx/session/redis"
	"github.com/gotomicro/ego/core/econf"
	"github.com/redis/go-redis/v9"
)

func InitSession(cmd redis.Cmdable) session.Provider {
	type Config struct {
		SessionKey string `json:"sessionKey"`
	}
	var cfg Config
	err := econf.UnmarshalKey("session", &cfg)
	if err != nil {
		panic(err)
	}
	sp := redis2.NewSessionProvider(cmd, cfg.SessionKey)
	return sp
}
