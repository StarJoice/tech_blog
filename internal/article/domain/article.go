//@Date 2024/12/9 16:00
//@Desc

package domain

type Article struct {
	Id int64
	// 对应作者
	Uid     int64
	Title   string
	Content string
	Ctime   int64
	Utime   int64
}
