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

// UserCollectBiz 收藏明细
type UserCollectBiz struct {
	Id    int64  `gorm:"primaryKey,autoIncrement"`
	Uid   int64  `gorm:"uniqueIndex:uid_biz_type_id"`
	BizId int64  `gorm:"uniqueIndex:uid_biz_type_id"`
	Biz   string `gorm:"type:varchar(128);uniqueIndex:uid_biz_type_id"`
	Utime int64
	Ctime int64
}

// Collection 收藏夹
type Collection struct {
	Id int64 `gorm:"primaryKey,autoIncrement"`
	// 在uid和name上创建唯一联合索引, 保证某一用户不会创建同名收藏夹
	Uid   int64  `gorm:"uniqueIndex:uid_name"`
	Name  string `gorm:"uniqueIndex:uid_name"`
	Ctime int64  `gorm:"autoCreateTime:milli"`
	Utime int64  `gorm:"autoUpdateTime:milli"`
}
