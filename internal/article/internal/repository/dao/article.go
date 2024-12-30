package dao

import (
	"context"
	"github.com/ego-component/egorm"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ArticleDao interface {
	Create(ctx context.Context, art Article) (int64, error)
	Update(ctx context.Context, art Article) error
	Count(ctx context.Context, uid int64) (int64, error)
	List(ctx context.Context, offset int, limit int, uid int64) ([]Article, error)
	Sync(ctx context.Context, art Article) (int64, error)
	PublishTotal(ctx context.Context) (int64, error)
	PubList(ctx context.Context, offset int, limit int) ([]PublishArticle, error)
	GetArtById(ctx context.Context, aid int64) (Article, error)
	GetPubArtById(ctx context.Context, aid int64) (PublishArticle, error)
	DeleteById(ctx context.Context, aid int64) error
}

type ArticleGormDao struct {
	db *egorm.Component
}

func NewArticleGormDao(db *egorm.Component) ArticleDao {
	return &ArticleGormDao{db: db}
}

func (dao *ArticleGormDao) DeleteById(ctx context.Context, aid int64) error {
	return dao.db.WithContext(ctx).Where("id=?", aid).Delete(&PublishArticle{}).Error
}

func (dao *ArticleGormDao) GetPubArtById(ctx context.Context, aid int64) (PublishArticle, error) {
	var art PublishArticle
	err := dao.db.WithContext(ctx).Where("id = ?", aid).First(&art).Error
	return art, err
}

func (dao *ArticleGormDao) GetArtById(ctx context.Context, aid int64) (Article, error) {
	var art Article
	err := dao.db.WithContext(ctx).Where("id=?", aid).First(&art).Error
	return art, err
}

func (dao *ArticleGormDao) PubList(ctx context.Context, offset int, limit int) ([]PublishArticle, error) {
	var arts []PublishArticle
	err := dao.db.WithContext(ctx).
		Select("id", "uid", "title", "abstract", "ctime").
		// 按id倒序排序
		Order("id desc").
		Model(&PublishArticle{}).
		Offset(offset).
		Limit(limit).
		Find(&arts).Error
	return arts, err
}

func (dao *ArticleGormDao) PublishTotal(ctx context.Context) (int64, error) {
	var res int64
	err := dao.db.WithContext(ctx).
		Model(&PublishArticle{}).
		Select("COUNT(id)").Count(&res).Error
	return res, err
}

func (dao *ArticleGormDao) Sync(ctx context.Context, art Article) (int64, error) {
	id := art.Id
	err := dao.db.Transaction(func(tx *gorm.DB) error {
		// 先判断这个要发表的帖子是之前的草稿存在于制作表还是说新作直接发表
		if art.Id == 0 {
			// 新作（保存到制作表）
			err := tx.WithContext(ctx).Create(&art).Error
			if err != nil {
				return err
			}
			id = art.Id
		} else {
			// 已有作品，先更新制作表
			err := dao.update(ctx, tx, art)
			if err != nil {
				return err
			}
		}
		// 直接保存到线上表
		publishArt := PublishArticle(art)
		// 如果保存值不包含主键，它将执行 Create，否则它将执行 Update (包含所有字段)
		// 这里实现upsert语义,想保持线上库和制作库的文章id保持一致
		//return tx.Save(&publishArt).Error
		return tx.Model(&PublishArticle{}).Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}}, // 根据 id 判断冲突
			UpdateAll: true,                          // 如果存在，更新所有字段
		}).Create(&publishArt).Error
	})
	return id, err
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
		// 查询列表 不显示文章具体内容，只显示摘要
		Select("id", "title", "abstract", "ctime").
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

func (dao *ArticleGormDao) update(ctx context.Context,
	tx *gorm.DB, art Article) error {
	return tx.WithContext(ctx).
		Model(&Article{}).Where("id = ?", art.Id).Updates(map[string]any{
		"title":   art.Title,
		"content": art.Content,
	}).Error
}
