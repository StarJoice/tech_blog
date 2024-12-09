//@Date 2024/12/9 16:01
//@Desc

package dao

import (
	"context"
	"github.com/ego-component/egorm"
)

type ArticleDao interface {
	Create(ctx context.Context, art Article) (int64, error)
	Update(ctx context.Context, art Article) error
	Count(ctx context.Context, uid int64) (int64, error)
	List(ctx context.Context, offset int, limit int, uid int64) ([]Article, error)
}

type ArticleGormDao struct {
	db *egorm.Component
}

func (dao *ArticleGormDao) Count(ctx context.Context, uid int64) (int64, error) {
	var count int64
	err := dao.db.WithContext(ctx).
		Model(&Article{}).
		Where("uid = ?", uid).
		Count(&count).Error
	return count, err
}

func (dao *ArticleGormDao) List(ctx context.Context,
	offset int, limit int, uid int64) ([]Article, error) {
	var articles []Article
	err := dao.db.WithContext(ctx).
		Where("uid = ?", uid).
		Select("id", "title", "content", "ctime").
		// 按创建时间倒序排序
		Order("ctime desc").
		Offset(offset).
		Limit(limit).
		Find(&articles).Error
	return articles, err
}

func (dao *ArticleGormDao) Create(ctx context.Context, art Article) (int64, error) {
	err := dao.db.WithContext(ctx).Create(&art).Error
	return art.Id, err
}

func (dao *ArticleGormDao) Update(ctx context.Context, art Article) error {
	return dao.db.WithContext(ctx).
		Model(&Article{}).
		Where("id = ?", art.Id).Updates(map[string]any{
		"title":   art.Title,
		"content": art.Content,
	}).Error
}

func NewArticleGormDao(db *egorm.Component) ArticleDao {
	return &ArticleGormDao{db: db}
}
