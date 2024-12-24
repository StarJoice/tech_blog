package domain

// Interactive 汇总表
type Interactive struct {
	// 标识资源（目前只有文章，后续或许会有评论的点赞，视频的点赞）
	Biz   string
	BizId int64
	// 阅读（观看）计数
	ViewCnt int64
	// 点赞计数
	LikeCnt int64
	// 收藏计数
	CollectCnt int64
}

// UserLikeBiz 点赞明细表
type UserLikeBiz struct {
}

// UserCollectionBiz 收藏明细表
type UserCollectionBiz struct {
}
