package dao

import "database/sql"

// Comment 使用树形结构实现评论多级嵌套
// 并设置删除策略为级联删除，即删除一条评论，数据库会自动删除的他的子评论
type Comment struct {
	Id  int64
	Uid int64
	// 这个代表的是你评论的对象是什么？
	// 比如说代表某个帖子，代表某个视频，代表某个图片
	Biz   string `gorm:"index:biz_type_id"`
	BizID int64  `gorm:"index:biz_type_id"`
	// 用 NULL 来表达没有父亲
	PId sql.NullInt64 `gorm:"index"`
	// 外键指向的也是同一张表 -- 级联删除
	ParentComment *Comment `gorm:"foreignKey:PId;associationForeignKey:Id;constraint:OnDelete:CASCADE"`
	// 引入 RootID 这个设计
	// 顶级评论的 ID
	// 主要是为了加载整棵评论的回复组成树
	RootID sql.NullInt64 `gorm:"index:root_ID_ctime"`
	// 评论的内容
	Content string
	Utime   int64 `gorm:"autoUpdateTime:milli"`
	Ctime   int64 `gorm:"index:root_ID_ctime;autoCreateTime:milli"`
}
