package dao

type User struct {
	Id       int64  `gorm:"primaryKey,autoIncrement"`
	Email    string `gorm:"unique"`
	Password string
	Nickname string
	Avatar   string
	AboutMe  string
	//HasPublishedArticles bool   // 是否发布过文章
	//Role                 string // 用户角色，比如 "reader" 或 "creator" 或者是 "admin" ???
	// 存储为毫秒时间戳
	Ctime int64 `gorm:"autoCreateTime:milli"` // 自动创建时间，单位为毫秒
	Utime int64 `gorm:"autoUpdateTime:milli"` // 自动更新时间，单位为毫秒
}
