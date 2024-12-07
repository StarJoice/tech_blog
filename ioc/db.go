//@Date 2024/12/5 01:16
//@Desc

package ioc

import (
	"github.com/ego-component/egorm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// InitDB 目前开发暂时使用gorm
func InitDB() *gorm.DB {
	DB, err := gorm.Open(mysql.Open("root:root@tcp(localhost:3307)/tech_blog"))
	if err != nil {
		panic("failed to connect database")
	}
	return DB
}

// InitDBv1 使用egorm
func InitDBv1() *egorm.Component {
	return egorm.Load("mysql").Build()
}
