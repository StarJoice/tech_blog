package domain

import "time"

const BizComment = "Comment"

type Comment struct {
	Id      int64  `json:"id"`
	Uid     int64  `json:"uid"`
	Biz     string `json:"biz"`
	BizId   int64  `json:"biz_id"`
	Content string `json:"content"`
	// 根评论
	RootComment *Comment `json:"rootComment"`
	// 父评论
	ParentComment *Comment `json:"parentComment"`
	// 子评论
	Children []Comment `json:"children"`
	Ctime    time.Time `json:"ctime"`
	Utime    time.Time `json:"utime"`
}
