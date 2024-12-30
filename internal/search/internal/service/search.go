package service

import (
	"context"
	"errors"
	"github.com/StarJoice/tech_blog/internal/search/internal/domain"
	"github.com/StarJoice/tech_blog/internal/search/internal/repository"
	"golang.org/x/sync/errgroup"
	"strings"
)

type SearchService interface {
	Search(ctx context.Context, offset, limit int, expr string) (*domain.SearchResult, error)
}
type Service struct {
	searchHandlers map[string]SearchHandler
}

func NewService(artRepo repository.ArticleRepo) SearchService {
	searchHandlers := map[string]SearchHandler{
		"article": NewArticleHandler(artRepo),
	}
	return &Service{searchHandlers: searchHandlers}
}

func (svc *Service) Search(ctx context.Context, offset, limit int, expr string) (*domain.SearchResult, error) {
	biz, keywords, err := svc.parseExpr(expr)
	if err != nil {
		return nil, err
	}
	var eg errgroup.Group
	res := &domain.SearchResult{}
	// 相当于在这里做一个分发， 如果用户发搜索请求到全部类目下，直接发all请求
	if biz == "all" {
		for _, handler := range svc.searchHandlers {
			bizHandler := handler
			eg.Go(func() error {
				return bizHandler.search(ctx, keywords, offset, limit, res)
			})
		}
		if err := eg.Wait(); err != nil {
			return nil, err
		}
	} else {
		bizHandler, ok := svc.searchHandlers[biz]
		if !ok {
			return nil, errors.New("无相关的业务处理方式")
		}
		err = bizHandler.search(ctx, keywords, offset, limit, res)
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}
func (svc *Service) parseExpr(expr string) (string, []domain.QueryMeta, error) {
	searchParams := strings.SplitN(expr, ":", 3)
	if len(searchParams) == 3 {
		typ := searchParams[0]
		if typ != "biz" {
			return "", nil, errors.New("参数错误")
		}
		biz := searchParams[1]
		keywords := searchParams[2]
		return biz, svc.getQueryMeta(keywords), nil
	}
	return "", nil, errors.New("参数错误")
}
func (svc *Service) getQueryMeta(keywords string) []domain.QueryMeta {
	keywordList := strings.Split(keywords, " ")
	metas := make([]domain.QueryMeta, 0, len(keywordList))
	for _, keyword := range keywordList {
		params := strings.Split(keyword, ":")
		if len(params) == 1 {
			metas = append(metas, domain.QueryMeta{
				Keyword: params[0],
				IsAll:   true,
			})
		}
		if len(params) == 2 {
			metas = append(metas, domain.QueryMeta{
				Keyword: params[1],
				Col:     params[0],
			})
		}
	}
	return metas
}
