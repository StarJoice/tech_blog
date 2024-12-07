//@Date 2024/12/5 00:42
//@Desc

package domain

type User struct {
	Id       int64
	Email    string
	Password string
	Nickname string
	Avatar   string
	AboutMe  string
	Ctime    int64
	Utime    int64
}
