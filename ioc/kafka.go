package ioc

import (
	"github.com/ecodeclub/mq-api"
	"github.com/ecodeclub/mq-api/kafka"
	"github.com/gotomicro/ego/core/econf"
)

func InitMq() (mq.MQ, error) {
	type Config struct {
		Network   string   `yaml:"network"`
		Addresses []string `yaml:"addresses"`
	}
	var cfg Config
	err := econf.UnmarshalKey("kafka", &cfg)
	if err != nil {
		return nil, err
	}
	qq, err := kafka.NewMQ(cfg.Network, cfg.Addresses)
	if err != nil {
		return nil, err
	}
	return qq, nil
}
