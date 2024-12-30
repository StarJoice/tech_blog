package dao

import (
	"github.com/ego-component/egorm"
)

func InitTable(db *egorm.Component) error {
	return db.AutoMigrate(&User{})
}

// TableName 实现tableName 接口，指定建表时表名
func (User) TableName() string {
	return "user"
}
