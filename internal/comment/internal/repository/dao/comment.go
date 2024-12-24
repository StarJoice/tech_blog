package dao

import (
	"context"
	"github.com/ego-component/egorm"
)

type CommentDao interface {
	Insert(ctx context.Context, c Comment) error
	// FindByBiz 查询某个资源的一级评论
	FindByBiz(ctx context.Context, biz string,
		bizID, minID, limit int64) ([]Comment, error)
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
func (dao *GormCommentDao) FindByBiz(ctx context.Context, biz string, bizID, minID, limit int64) ([]Comment, error) {
	var res []Comment
	err := dao.db.WithContext(ctx).
		// 只要顶级评论 所以限制pid=null
		Where("biz = ? AND biz_ID = ? AND id < ? AND pid IS NULL", biz, bizID, minID).
		Limit(int(limit)).Find(&res).
		Error
	return res, err
}
