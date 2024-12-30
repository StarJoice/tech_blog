package dao

import (
	"context"
	"github.com/olivere/elastic/v7"
)

type AnyDao interface {
	Input(ctx context.Context, index string, docID string, data string) error
}

type anyEsDao struct {
	client *elastic.Client
}

func NewAnyDao(client *elastic.Client) AnyDao {
	return &anyEsDao{
		client: client,
	}
}
func (a *anyEsDao) Input(ctx context.Context,
	index string, docID string, data string) error {
	_, err := a.client.Index().
		Index(index).
		Id(docID).
		BodyJson(data).Do(ctx)
	return err
}
