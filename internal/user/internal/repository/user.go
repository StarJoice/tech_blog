//@Date 2024/12/5 00:46
//@Desc

package repository

import (
	"context"
	"github.com/StarJoice/tech_blog/internal/user/internal/domain"
	"github.com/StarJoice/tech_blog/internal/user/internal/repository/cache"
	"github.com/StarJoice/tech_blog/internal/user/internal/repository/dao"
)

var (
	ErrDuplicateEmail = dao.ErrDuplicateEmail
	ErrUserNotFound   = dao.ErrRecordNotFound
)

//go:generate mockgen -source=./user.go -package=repomocks -destination=mocks/user.mock.go UserRepository
type UserRepository interface {
	Create(ctx context.Context, user domain.User) error
	FindByEmail(ctx context.Context, email string) (domain.User, error)
	FindById(ctx context.Context, uid int64) (domain.User, error)
	Update(ctx context.Context, user domain.User) error
}

type UserCacheRepository struct {
	dao dao.UserDao
	// 暂时组装起来，但不使用
	cache cache.UserCache
}

func NewUserCacheRepository(dao dao.UserDao) UserRepository {
	return &UserCacheRepository{dao: dao}
}
func (repo *UserCacheRepository) FindById(ctx context.Context, uid int64) (domain.User, error) {
	// todo 后续或许可以接入缓存
	u, err := repo.dao.FindById(ctx, uid)
	return repo.toDomain(u), err
}

func (repo *UserCacheRepository) Update(ctx context.Context, user domain.User) error {
	return repo.dao.UpdateNonZeroFields(ctx, repo.domainToEntity(user))
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
	return repo.toDomain(u), nil
}

func (repo *UserCacheRepository) domainToEntity(u domain.User) dao.User {
	// 暂时只有更新接口会使用，所以不引入非敏感字段以外的字段
	return dao.User{
		Id:       u.Id,
		Nickname: u.Nickname,
		Avatar:   u.Avatar,
		AboutMe:  u.AboutMe,
	}
}

func (repo *UserCacheRepository) toDomain(u dao.User) domain.User {
	return domain.User{
		Id:       u.Id,
		Email:    u.Email,
		Password: u.Password,
		Nickname: u.Nickname,
		Avatar:   u.Avatar,
		AboutMe:  u.AboutMe,
	}
}
