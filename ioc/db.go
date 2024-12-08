//@Date 2024/12/5 01:16
//@Desc

package ioc

import (
	"github.com/ego-component/egorm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// InitDBv1 开发使用gorm
func InitDBv1() *gorm.DB {
	DB, err := gorm.Open(mysql.Open("root:root@tcp(localhost:3307)/tech_blog"))
	if err != nil {
		panic("failed to connect database")
	}
	return DB
}

// InitDB InitDBv1 使用egorm todo 要引入一定的重试策略
func InitDB() *egorm.Component {
	return egorm.Load("mysql").Build()
}
