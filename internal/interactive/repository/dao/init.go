package dao

import "github.com/ego-component/egorm"

func InitTable(db *egorm.Component) error {
	return db.AutoMigrate(&Interactive{}, &UserLikeBiz{}, &UserCollectBiz{})
}
func (Interactive) TableName() string {
	return "interactive"
}
func (UserLikeBiz) TableName() string {
	return "userLikeBiz"
}
func (UserCollectBiz) TableName() string {
	return "userCollectBiz"
}
