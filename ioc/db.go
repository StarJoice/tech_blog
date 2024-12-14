//@Date 2024/12/5 01:16
//@Desc

package ioc

import (
	"github.com/ego-component/egorm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// InitDB 初始化mysql DB
func InitDB() *egorm.Component {
	DB, err := gorm.Open(mysql.Open("root:root@tcp(localhost:3307)/tech_blog"))
	if err != nil {
		panic("failed to connect database")
	}
	return DB
}
