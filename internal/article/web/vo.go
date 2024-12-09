//@Date 2024/12/9 20:45
//@Desc

package web

import "github.com/StarJoice/tech_blog/internal/article/domain"

type Article struct {
	Id      int64  `json:"id,omitempty"`
	Uid     int64  `json:"uid,omitempty"`
	Title   string `json:"title,omitempty" validate:"required"`
	Content string `json:"content,omitempty"`
	Ctime   int64  `json:"ctime,omitempty"`
}

type SaveReq struct {
	Id      int64  `json:"id,omitempty"`
	Uid     int64  `json:"uid,omitempty"`
	Title   string `json:"title,omitempty" binding:"required"`
	Content string `json:"content,omitempty"`
}

func (a SaveReq) toDomain() domain.Article {
	return domain.Article{
		Id:      a.Id,
		Uid:     a.Uid,
		Title:   a.Title,
		Content: a.Content,
	}
}

type ArtsList struct {
	Arts  []Article `json:"arts,omitempty"`
	Total int64     `json:"total,omitempty"`
}
