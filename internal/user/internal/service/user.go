package service

import (
	"context"
	"errors"
	"github.com/StarJoice/tech_blog/internal/user/internal/domain"
	"github.com/StarJoice/tech_blog/internal/user/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrDuplicateEmail        = repository.ErrDuplicateEmail
	ErrInvalidUserOrPassword = errors.New("用户不存在或密码不正确")
	ErrInvalidPassword       = repository.ErrInvalidPassword
)

//go:generate mockgen -source=./user.go -package=svcmocks -destination=mocks/user.mock.go UserService
type UserService interface {
	Signup(ctx context.Context, user domain.User) error
	Login(ctx context.Context, email string, password string) (domain.User, error)
	Profile(ctx context.Context, uid int64) (domain.User, error)
	// UpdateNonSensitiveInfo 更新用户信息下的非敏感字段（就是指头像昵称等等...）
	UpdateNonSensitiveInfo(ctx context.Context, user domain.User) error
	UpdatePassword(ctx context.Context, uid int64, oldPwd string, newPwd string) error
	// GetByID 提供给文章模块来查询文章作者
	GetByID(ctx context.Context, id int64) (domain.User, error)
}

type UserSvc struct {
	repo repository.UserRepository
}

func NewUserSvc(repo repository.UserRepository) UserService {
	return &UserSvc{repo: repo}
}

func (svc *UserSvc) Signup(ctx context.Context, user domain.User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hash)
	return svc.repo.Create(ctx, user)
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

func (svc *UserSvc) Profile(ctx context.Context, uid int64) (domain.User, error) {
	return svc.repo.FindById(ctx, uid)
}

func (svc *UserSvc) UpdateNonSensitiveInfo(ctx context.Context, user domain.User) error {
	return svc.repo.Update(ctx, user)
}

func (svc *UserSvc) UpdatePassword(ctx context.Context, uid int64, oldPwd string, newPwd string) error {
	return svc.repo.UpdatePassword(ctx, uid, oldPwd, newPwd)
}

func (svc *UserSvc) GetByID(ctx context.Context, id int64) (domain.User, error) {
	return svc.repo.GetByID(ctx, id)
}
