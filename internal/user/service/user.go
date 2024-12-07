//@Date 2024/12/5 00:42
//@Desc

package service

import (
	"context"
	"errors"
	"github.com/StarJoice/tech_blog/internal/user/domain"
	"github.com/StarJoice/tech_blog/internal/user/repository"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrDuplicateEmail        = repository.ErrDuplicateEmail
	ErrInvalidUserOrPassword = errors.New("用户不存在或密码不正确")
)

type UserService interface {
	Signup(ctx context.Context, user domain.User) error
	Login(ctx context.Context, email string, password string) (domain.User, error)
}

type UserSvc struct {
	repo repository.UserRepository
}

func (svc *UserSvc) Login(ctx context.Context, email string, password string) (domain.User, error) {
	// 查找用户
	u, err := svc.repo.FindByEmail(ctx, email)
	if errors.Is(err, repository.ErrUserNotFound) {
		return domain.User{}, ErrInvalidUserOrPassword
	}
	if err != nil {
		return domain.User{}, err
	}
	// 验证密码是否正确
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return domain.User{}, ErrInvalidUserOrPassword
	}
	return u, nil
}

func (svc *UserSvc) Signup(ctx context.Context, user domain.User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hash)
	return svc.repo.Create(ctx, user)
}

func NewUserSvc(repo repository.UserRepository) UserService {
	return &UserSvc{repo: repo}
}
