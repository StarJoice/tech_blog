//@Date 2024/12/10 00:06
//@Desc

package dao

import "github.com/ego-component/egorm"

func InitTable(db *egorm.Component) error {
	return db.AutoMigrate(&Label{})
}
