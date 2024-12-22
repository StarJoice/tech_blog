package repository

import "github.com/StarJoice/tech_blog/internal/label/repository/dao"

type Repository interface {
}
type LabelRepository struct {
	dao dao.LabelDao
}

func NewLabelRepository(dao dao.LabelDao) Repository {
	return &LabelRepository{dao: dao}
}
