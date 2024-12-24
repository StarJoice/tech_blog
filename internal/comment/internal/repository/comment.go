package repository

import "github.com/StarJoice/tech_blog/internal/comment/internal/repository/dao"

type Repository interface {
}
type CommentRepository struct {
	dao dao.CommentDao
}
