//@Date 2024/12/5 00:52
//@Desc

package dao

import (
	"gorm.io/gorm"
)

func InitTable(db *gorm.DB) error {
	return db.AutoMigrate(&User{})
}
