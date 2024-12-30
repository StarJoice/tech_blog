package dao

import "github.com/ego-component/egorm"

func InitTable(db *egorm.Component) error {
	return db.AutoMigrate(&Article{}, &PublishArticle{})
}

// TableName 设置表名
func (Article) TableName() string {
	return "articles"
}
func (PublishArticle) TableName() string {
	return "publish_articles"
}
