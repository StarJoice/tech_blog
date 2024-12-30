package web

type searchReq struct {
	Offset   int    `json:"offset"`   // 偏移量
	Limit    int    `json:"limit"`    // 一页几条数据
	Keywords string `json:"keywords"` // 搜索关键词
}
