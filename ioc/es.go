package ioc

import (
	"fmt"
	"github.com/gotomicro/ego/core/econf"
	"github.com/olivere/elastic/v7"
	"time"
)

func InitES() *elastic.Client {
	type Config struct {
		Url   string `yaml:"url"`
		Sniff bool   `yaml:"sniff"`
	}
	var cfg Config
	err := econf.UnmarshalKey("es", &cfg)
	if err != nil {
		panic(fmt.Errorf("读取 ES 配置失败 %w", err))
	}
	const timeout = 10 * time.Second
	opts := []elastic.ClientOptionFunc{
		elastic.SetURL(cfg.Url),
		elastic.SetSniff(cfg.Sniff),
		elastic.SetHealthcheckTimeoutStartup(timeout),
	}
	client, err := elastic.NewClient(opts...)
	if err != nil {
		panic(err)
	}
	return client
}
