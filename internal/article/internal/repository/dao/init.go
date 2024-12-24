package dao

import "github.com/ego-component/egorm"

func InitTable(db *egorm.Component) error {
	return db.AutoMigrate(&Article{}, &PublishArticle{})
}
