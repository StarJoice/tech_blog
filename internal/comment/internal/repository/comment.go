package repository

import (
	"context"
	"database/sql"
	"github.com/StarJoice/tech_blog/internal/comment/internal/domain"
	"github.com/StarJoice/tech_blog/internal/comment/internal/repository/dao"
)

type Repository interface {
	CreateComment(ctx context.Context, comment domain.Comment) error
}
type CachedCommentRepository struct {
	dao dao.CommentDao
}

func NewCommentRepository(dao dao.CommentDao) Repository {
	return &CachedCommentRepository{dao: dao}
}

func (repo *CachedCommentRepository) CreateComment(ctx context.Context, comment domain.Comment) error {
	return repo.dao.Insert(ctx, repo.toEntity(comment))
}

func (repo *CachedCommentRepository) toEntity(comment domain.Comment) dao.Comment {
	daoComment := dao.Comment{
		Id:      comment.Id,
		Uid:     comment.Uid,
		Biz:     comment.Biz,
		BizID:   comment.BizId,
		Content: comment.Content,
	}
	if comment.ParentComment != nil {
		daoComment.PId = sql.NullInt64{
			Int64: comment.ParentComment.Id,
			Valid: true,
		}
	}
	if comment.RootComment != nil {
		daoComment.RootID = sql.NullInt64{
			Int64: comment.RootComment.Id,
			Valid: true,
		}
	}
	return daoComment
}
