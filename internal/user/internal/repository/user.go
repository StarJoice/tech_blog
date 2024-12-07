//@Date 2024/12/5 00:46
//@Desc

package repository

import (
	"context"
	"github.com/StarJoice/tech_blog/internal/user/internal/domain"
	"github.com/StarJoice/tech_blog/internal/user/internal/repository/dao"
)

var (
	ErrDuplicateEmail = dao.ErrDuplicateEmail
	ErrUserNotFound   = dao.ErrRecordNotFound
)

type UserRepository interface {
	Create(ctx context.Context, user domain.User) error
	FindByEmail(ctx context.Context, email string) (domain.User, error)
}

type UserCacheRepository struct {
	dao dao.UserDao
	//cache cache.UserCache
}

func NewUserCacheRepository(dao dao.UserDao) UserRepository {
	return &UserCacheRepository{dao: dao}
}

func (repo *UserCacheRepository) Create(ctx context.Context, user domain.User) error {
	return repo.dao.Insert(ctx, dao.User{
		Email:    user.Email,
		Password: user.Password,
	})
}
func (repo *UserCacheRepository) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	u, err := repo.dao.FindByEmail(ctx, email)
	if err != nil {
		return domain.User{}, err
	}
	return repo.ToDomain(u), nil
}

func (repo *UserCacheRepository) ToDomain(u dao.User) domain.User {
	return domain.User{
		Id:       u.Id,
		Email:    u.Email,
		Password: u.Password,
		Nickname: u.Nickname,
		Avatar:   u.Avatar,
		AboutMe:  u.AboutMe,
	}
}
