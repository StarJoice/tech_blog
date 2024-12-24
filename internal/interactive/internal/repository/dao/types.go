package dao

// Interactive 互动汇总表
type Interactive struct {
	Id    int64  `gorm:"primaryKey,autoIncrement"`
	BizId int64  `gorm:"uniqueIndex:biz_type_id"`
	Biz   string `gorm:"type:varchar(128);uniqueIndex:biz_type_id"`

	ViewCnt    int64
	LikeCnt    int64
	CollectCnt int64
	Ctime      int64 `gorm:"autoCreateTime:milli"` // 自动创建时间，单位为毫秒
	Utime      int64 `gorm:"autoUpdateTime:milli"` // 自动更新时间，单位为毫秒
}

// UserLikeBiz 用户点赞明细表
type UserLikeBiz struct {
	Id    int64  `gorm:"primaryKey,autoIncrement"`
	Uid   int64  `gorm:"uniqueIndex:uid_biz_type_id"`
	BizId int64  `gorm:"uniqueIndex:uid_biz_type_id"`
	Biz   string `gorm:"type:varchar(128);uniqueIndex:biz_type_id"`

	Ctime int64 `gorm:"autoCreateTime:milli"`
	Utime int64 `gorm:"autoUpdateTime:milli"`
}

// UserCollectionBiz 收藏明细
type UserCollectionBiz struct {
	Id    int64  `gorm:"primaryKey,autoIncrement"`
	Uid   int64  `gorm:"uniqueIndex:uid_biz_type_id"`
	BizId int64  `gorm:"uniqueIndex:uid_biz_type_id"`
	Biz   string `gorm:"type:varchar(128);uniqueIndex:uid_biz_type_id"`
	// 收藏夹id
	Cid   int64 `gorm:"index;not null;default(0)"`
	Utime int64 `gorm:"autoUpdateTime:milli"`
	Ctime int64 `gorm:"autoCreateTime:milli"`
}

// Collection 收藏夹
type Collection struct {
	Id int64 `gorm:"primaryKey,autoIncrement"`
	// 在 Uid 和 Name 上创建唯一索引，确保用户不会创建同名收藏夹
	Uid   int64  `gorm:"uniqueIndex:uid_name"`
	Name  string `gorm:"type:varchar(256);uniqueIndex:uid_name"`
	Ctime int64  `gorm:"autoCreateTime:milli"`
	Utime int64  `gorm:"autoUpdateTime:milli"`
}
