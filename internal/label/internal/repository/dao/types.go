package dao

type Label struct {
	Id    int64 `gorm:"primary_key;AUTO_INCREMENT"`
	Uid   int64
	Name  string
	Ctime int64 `gorm:"autoCreateTime:milli"`       // 自动创建时间，单位为毫秒
	Utime int64 `gorm:"autoUpdateTime:milli;index"` // 自动更新时间，单位为毫秒
}

func (Label) TableName() string {
	return "label"
}
