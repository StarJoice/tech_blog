package dao

// Article 制作库
type Article struct {
	Id int64 `gorm:"primaryKey;autoIncrement"`
	// 对应作者
	Uid      int64  `gorm:"index"`
	Title    string `gorm:"not null"`
	Content  string `gorm:"type:text"`
	Abstract string `gorm:"type:text"`
	// 存储为毫秒时间戳
	Ctime int64 `gorm:"autoCreateTime:milli"`       // 自动创建时间，单位为毫秒
	Utime int64 `gorm:"autoUpdateTime:milli;index"` // 自动更新时间，单位为毫秒
}

type PublishArticle Article
