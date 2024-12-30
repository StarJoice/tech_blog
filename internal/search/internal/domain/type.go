package domain

import (
	"sync"
	"time"
)

type User struct {
	Id       int64
	Email    string
	Nickname string
}
type Article struct {
	Id int64
	// 对应作者
	Uid int64
	// 文章标题
	Title string
	// 文章内容
	Content string
	// 摘要
	Abstract string
	Ctime    time.Time
	Utime    time.Time
}

type SearchResult struct {
	mu       sync.RWMutex
	Articles []Article
}

func (s *SearchResult) SetArticles(articles []Article) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Articles = articles
}
