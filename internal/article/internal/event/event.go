package event

import (
	"encoding/json"
	"github.com/StarJoice/tech_blog/internal/article/internal/domain"
)

type ArticleEvent struct {
	Biz   string `json:"biz"`
	BizID int    `json:"bizId"`
	Data  string `json:"data"`
}

type Article struct {
	Id int64 `json:"id"`
	// 对应作者
	Uid int64 `json:"uid"`
	// 文章标题
	Title string `json:"title"`
	// 文章内容
	Content string `json:"content"`
	// 摘要
	Abstract string `json:"abstract"`
	Ctime    int64  `json:"ctime"`
	Utime    int64  `json:"utime"`
}

func newArticle(art domain.Article) Article {
	return Article{
		Id:       art.Id,
		Uid:      art.Uid,
		Title:    art.Title,
		Content:  art.Content,
		Abstract: art.Abstract,
		Ctime:    art.Ctime.UnixMilli(),
		Utime:    art.Utime.UnixMilli(),
	}
}

// NewArticleEvent 创建一个文章事件，封装文章的业务信息和数据
// 这个函数接收一个 domain.Article 结构体，将其转换为 Article 结构体，并
// 序列化为 JSON 字符串后返回一个 ArticleEvent 对象。
// ArticleEvent 是文章业务事件的标准结构，便于事件的发布和处理。
func NewArticleEvent(art domain.Article) ArticleEvent {
	// 将文章结构体转换为 JSON 格式
	data, _ := json.Marshal(newArticle(art))
	return ArticleEvent{
		Biz:   "article",    // 业务类型为文章
		BizID: int(art.Id),  // 文章的 ID
		Data:  string(data), // 文章的具体数据（JSON 格式）
	}
}
