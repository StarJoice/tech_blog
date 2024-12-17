package domain

import "time"

type Article struct {
	Id int64
	// 对应作者
	Uid int64
	// 文章标题
	Title string
	// 文章内容
	Content string
	// 摘要
	Abstract string
	Ctime    time.Time
	Utime    time.Time
}
