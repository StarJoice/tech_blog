//@Date 2024/12/5 01:16
//@Desc

package ioc

import (
	"github.com/ego-component/egorm"
)

// InitDB 初始化mysql DB
func InitDB() *egorm.Component {
	return egorm.Load("mysql").Build()
}
