package web

import "github.com/StarJoice/tech_blog/internal/comment/internal/domain"

type createCommentReq struct {
	Uid           int64           `json:"uid"`
	Biz           string          `json:"biz"`
	BizId         int64           `json:"biz_id"`
	Content       string          `json:"content"`
	ParentComment *domain.Comment `json:"parent_comment"`
	RootComment   *domain.Comment `json:"root_comment"`
}
