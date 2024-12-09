//@Date 2024/12/9 16:09
//@Desc

package dao

type User struct {
	Id       int64  `gorm:"primaryKey,autoIncrement"`
	Email    string `gorm:"unique"`
	Password string
	Nickname string
	Avatar   string
	AboutMe  string
	// 存储为毫秒时间戳
	Ctime int64 `gorm:"autoCreateTime:milli"` // 自动创建时间，单位为毫秒
	Utime int64 `gorm:"autoUpdateTime:milli"` // 自动更新时间，单位为毫秒
}

// TableName 实现tableName 接口，指定建表时表名
func (User) TableName() string {
	return "user"
}
