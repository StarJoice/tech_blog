package web

type likeReq struct {
	Biz   string `json:"biz"`
	BizId int64  `json:"biz_id"`
	// true => 点赞
	// false => 取消点赞
	Liked bool `json:"liked"`
}
