package dao

import (
	"context"
	"encoding/json"
	"github.com/StarJoice/tech_blog/internal/search/internal/domain"
	"github.com/olivere/elastic/v7"
)

const ArticleIndexName = "article_index"

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

const (
	ArticleTitle   = 30
	ArticleContent = 5
	ArticleLabel   = 20
)

type ArticleDao interface {
	SearchArticle(ctx context.Context, offset,
		limit int, queryMetas []domain.QueryMeta) ([]Article, error)
}

type ArticleEsDao struct {
	client *elastic.Client
	metas  map[string]Col
}

func (es *ArticleEsDao) SearchArticle(ctx context.Context, offset,
	limit int, queryMetas []domain.QueryMeta) ([]Article, error) {
	query := elastic.NewBoolQuery().Must(
		elastic.NewBoolQuery().Should(buildCols(es.metas, queryMetas)...))
	resp, err := es.client.Search(ArticleIndexName).
		From(offset).
		Size(limit).
		Query(query).
		Do(ctx)
	if err != nil {
		return nil, err
	}
	res := make([]Article, 0, len(resp.Hits.Hits))
	for _, hit := range resp.Hits.Hits {
		var art Article
		err = json.Unmarshal(hit.Source, &art)
		if err != nil {
			return nil, err
		}
		res = append(res, art)
	}
	return res, nil
}

func NewArticleEsDao(client *elastic.Client) *ArticleEsDao {
	return &ArticleEsDao{
		client: client,
		metas: map[string]Col{
			"title": {
				Name:  "title",
				Boost: ArticleTitle,
			},
			"content": {
				Name:  "content",
				Boost: ArticleContent,
			},
		},
	}
}
