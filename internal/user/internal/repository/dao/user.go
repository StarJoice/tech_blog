package dao

import (
	"context"
	"errors"
	"github.com/ego-component/egorm"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

var (
	ErrDuplicateEmail = errors.New("用户已存在")
	ErrRecordNotFound = gorm.ErrRecordNotFound
)

//go:generate mockgen -source=./user.go -package=daomocks -destination=mocks/user.mock.go UserDao
type UserDao interface {
	Insert(ctx context.Context, user User) error
	FindByEmail(ctx context.Context, email string) (User, error)
	FindById(ctx context.Context, uid int64) (User, error)
	UpdateNonZeroFields(ctx context.Context, u User) error
}

type UserGormDao struct {
	db *egorm.Component
}

func NewUserGormDao(db *egorm.Component) UserDao {
	return &UserGormDao{db: db}
}

func (dao *UserGormDao) UpdateNonZeroFields(ctx context.Context, u User) error {
	return dao.db.WithContext(ctx).Updates(&u).Error
}

func (dao *UserGormDao) FindById(ctx context.Context, uid int64) (User, error) {
	var u User
	err := dao.db.WithContext(ctx).First(&u, "Id=?", uid).Error
	return u, err
}

func (dao *UserGormDao) FindByEmail(ctx context.Context, email string) (User, error) {
	var user User
	err := dao.db.WithContext(ctx).First(&user, "email = ?", email).Error
	return user, err
}

func (dao *UserGormDao) Insert(ctx context.Context, user User) error {
	err := dao.db.WithContext(ctx).Create(&user).Error
	var mysqlError *mysql.MySQLError
	if errors.As(err, &mysqlError) {
		if mysqlError.Number == 1062 {
			return ErrDuplicateEmail
		}
	}
	return err
}
