package dao

import "github.com/ego-component/egorm"

type LabelDao interface{}
type LabelGormDao struct {
	db *egorm.Component
}

func NewLabelGormDao(db *egorm.Component) LabelDao {
	return &LabelGormDao{db: db}
}
