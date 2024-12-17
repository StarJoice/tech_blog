package ioc

import (
	"github.com/ego-component/egorm"
	"github.com/gotomicro/ego/core/econf"
)

var db *egorm.Component

func InitDB() *egorm.Component {
	if db != nil {
		return db
	}
	econf.Set("mysql", map[string]string{"dsn": "root:root@tcp(localhost:3307)/tech_blog?charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=True&loc=Local&timeout=1s&readTimeout=3s&writeTimeout=3s"})
	db = egorm.Load("mysql").Build()
	return db
}
