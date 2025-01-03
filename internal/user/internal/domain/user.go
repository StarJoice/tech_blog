package domain

import "time"

const BizUser = "User"

// User 用户模型
// @Description 用户信息结构体
type User struct {
	Id       int64
	Email    string
	Password string
	Nickname string
	Avatar   string
	AboutMe  string
	// todo
	//HasPublishedArticles bool   // 是否发布过文章
	//Role                 string // 用户角色，比如 "reader" 或 "creator" 或者是 "admin" ???
	Ctime time.Time
	Utime time.Time
}
