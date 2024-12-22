package dao

import (
	"context"
	"github.com/ego-component/egorm"
)

type CommentDao interface {
	Insert(ctx context.Context, c Comment) error
}
type GormCommentDao struct {
	db *egorm.Component
}

func NewCommentDao(db *egorm.Component) CommentDao {
	return &GormCommentDao{db: db}
}
func (dao *GormCommentDao) Insert(ctx context.Context, c Comment) error {
	return dao.db.WithContext(ctx).Create(&c).Error
}
