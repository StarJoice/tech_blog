package dao

import "github.com/ego-component/egorm"

func InitTable(db *egorm.Component) error {
	return db.AutoMigrate(&Interactive{}, &UserLikeBiz{}, &UserCollectionBiz{}, &Collection{})
}
func (Interactive) TableName() string {
	return "interactive"
}
func (UserLikeBiz) TableName() string {
	return "userLikeBiz"
}
func (UserCollectionBiz) TableName() string {
	return "userCollectBiz"
}
func (Collection) TableName() string {
	return "collection"
}
