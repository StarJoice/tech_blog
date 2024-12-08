//@Date 2024/12/5 00:52
//@Desc

package dao

import (
	"github.com/ego-component/egorm"
)

func InitTable(db *egorm.Component) error {
	return db.AutoMigrate(&User{})
}
