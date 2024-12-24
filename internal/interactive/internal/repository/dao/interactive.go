package dao

import (
	"context"
	"errors"
	"github.com/ego-component/egorm"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var ErrRecordNotFound = gorm.ErrRecordNotFound

type InteractiveDao interface {
	// IncrViewCnt 增加阅读数量
	IncrViewCnt(ctx context.Context, biz string, bizId int64) error
	// LikeToggle 点赞或者取消点赞
	LikeToggle(ctx context.Context, uid int64, biz string, id int64) error
	// GetLikeInfo 获取点赞信息（是否点赞）
	GetLikeInfo(ctx context.Context, uid int64, biz string, id int64) (UserLikeBiz, error)
	// Get 获取某个资源的互动明细
	Get(ctx context.Context, biz string, id int64) (Interactive, error)
	GetByIds(ctx context.Context, biz string, ids []int64) ([]Interactive, error)
	// CollectToggle 收藏或者取消收藏
	CollectToggle(ctx context.Context, c UserCollectionBiz) error
}
type InteractiveGormDao struct {
	db *egorm.Component
}

func NewInteractiveGormDao(db *egorm.Component) InteractiveDao {
	return &InteractiveGormDao{db: db}
}

func (dao *InteractiveGormDao) IncrViewCnt(ctx context.Context, biz string, bizId int64) error {
	return dao.db.WithContext(ctx).Clauses(clause.OnConflict{
		// 如果有数据，直接计数加一
		DoUpdates: clause.Assignments(map[string]any{
			"view_cnt": gorm.Expr("view_cnt + ?", 1),
		}),
	}). // 此处是 当数据库中，没有产生主键冲突的话，直接新建一条数据
		Create(&Interactive{
			BizId:   bizId,
			Biz:     biz,
			ViewCnt: 1,
		}).Error
}

func (dao *InteractiveGormDao) LikeToggle(ctx context.Context, uid int64, biz string, id int64) error {
	return dao.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.
			Where("uid = ? AND biz_id = ? AND id = ?", uid, biz, id).
			First(&biz).Error
		switch {
		case err == nil:
			// 如果能查询到数据，那么要取消点赞
			return dao.deleteLikedInfo(tx, uid, biz, id)
		case errors.Is(err, gorm.ErrRecordNotFound):
			// 如果查询不到数据，就是要点赞，就是新建数据
			return dao.insertLikeInfo(tx, uid, biz, id)
		default:
			return err
		}
	})
}

// 新建点赞数据
func (dao *InteractiveGormDao) insertLikeInfo(tx *gorm.DB,
	uid int64, biz string, id int64) error {
	err := tx.Create(&UserLikeBiz{
		Uid:   uid,
		BizId: id,
		Biz:   biz,
	}).Error
	if err != nil {
		return err
	}
	return tx.Clauses(clause.OnConflict{
		DoUpdates: clause.Assignments(map[string]any{
			"like_cnt": gorm.Expr("`like_cnt` + 1"),
		}),
	}).Create(&Interactive{
		Biz:     biz,
		BizId:   id,
		LikeCnt: 1,
	}).Error
}

// 删除点赞数据
func (dao *InteractiveGormDao) deleteLikedInfo(tx *gorm.DB,
	uid int64, biz string, id int64) error {
	res := tx.Model(&UserLikeBiz{}).
		Where("uid = ? AND biz_id = ? AND id = ?", uid, biz, id).
		Delete(&UserLikeBiz{})
	if res.Error != nil {
		return res.Error
	}
	// 如果 没有删除任何数据，直接返回
	if res.RowsAffected < 1 {
		return nil
	}
	return tx.Model(&Interactive{}).
		Where("biz =? AND biz_id=?", biz, id).
		Updates(map[string]any{
			"like_cnt": gorm.Expr("`like_cnt` - 1"),
		}).Error
}
func (dao *InteractiveGormDao) GetLikeInfo(ctx context.Context, uid int64, biz string, id int64) (UserLikeBiz, error) {
	var res UserLikeBiz
	err := dao.db.WithContext(ctx).
		Where("biz = ? AND biz_id = ? AND uid = ?", biz, id, uid).Find(&res).Error
	return res, err
}
func (dao *InteractiveGormDao) Get(ctx context.Context, biz string, id int64) (Interactive, error) {
	var res Interactive
	err := dao.db.WithContext(ctx).
		Where("biz = ? AND biz_id = ?", biz, id).Find(&res).Error
	return res, err
}
func (dao *InteractiveGormDao) GetByIds(ctx context.Context, biz string, ids []int64) ([]Interactive, error) {
	var res []Interactive
	err := dao.db.WithContext(ctx).
		Where("biz = ? AND biz_id IN ?", biz, ids).
		Order("biz_id desc").
		Find(&res).Error
	return res, err
}
func (dao *InteractiveGormDao) CollectToggle(ctx context.Context, c UserCollectionBiz) error {
	return dao.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.
			Where("biz = ? AND biz_id = ? AND uid = ?", c.Biz, c.BizId, c.Uid).
			First(&UserCollectionBiz{}).Error
		switch {
		case err == nil:
			// 如果有数据，切换到未收藏状态
			return dao.deleteCollectionInfo(tx, c.Biz, c.BizId, c.Uid)
		case errors.Is(err, gorm.ErrRecordNotFound):
			// 没有数据，则添加数据
			return dao.insertCollectionInfo(tx, c)
		default:
			return err
		}
	})
}

func (dao *InteractiveGormDao) insertCollectionInfo(tx *gorm.DB,
	c UserCollectionBiz) error {
	err := tx.Create(&c).Error
	if err != nil {
		return err
	}
	return tx.Clauses(clause.OnConflict{
		DoUpdates: clause.Assignments(map[string]any{
			"collect_cnt": gorm.Expr("`collect_cnt` + 1"),
		}),
	}).Create(&Interactive{
		BizId:      c.BizId,
		Biz:        c.Biz,
		CollectCnt: 1,
	}).Error
}

func (dao *InteractiveGormDao) deleteCollectionInfo(tx *gorm.DB,
	biz string, id int64, uid int64) error {
	res := tx.Model(&UserCollectionBiz{}).
		Where("biz = ? AND biz_id = ? AND uid = ?", biz, id, uid)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected < 1 {
		return nil
	}
	return tx.Model(&Interactive{}).
		Where("biz = ? AND biz_id = ?", biz, id).
		Updates(map[string]any{
			"collect_cnt": gorm.Expr("`collect_cnt` - 1"),
		}).Error
}
