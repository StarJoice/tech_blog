package repository

import (
	"context"
	"errors"
	"github.com/StarJoice/tech_blog/internal/user/internal/domain"
	"github.com/StarJoice/tech_blog/internal/user/internal/repository/cache"
	"github.com/StarJoice/tech_blog/internal/user/internal/repository/dao"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var (
	ErrDuplicateEmail  = dao.ErrDuplicateEmail
	ErrUserNotFound    = dao.ErrRecordNotFound
	ErrInvalidPassword = errors.New("密码错误")
)

//go:generate mockgen -source=./user.go -package=repomocks -destination=mocks/user.mock.go UserRepository
type UserRepository interface {
	Create(ctx context.Context, user domain.User) error
	FindByEmail(ctx context.Context, email string) (domain.User, error)
	FindById(ctx context.Context, uid int64) (domain.User, error)
	Update(ctx context.Context, user domain.User) error
	UpdatePassword(ctx context.Context, uid int64, oldPwd string, newPwd string) error
	GetByID(ctx context.Context, id int64) (domain.User, error)
}

type UserCacheRepository struct {
	dao   dao.UserDao
	cache cache.UserCache
}

func NewUserCacheRepository(dao dao.UserDao, userCache cache.UserCache) UserRepository {
	return &UserCacheRepository{dao: dao, cache: userCache}
}

func (repo *UserCacheRepository) GetByID(ctx context.Context, id int64) (domain.User, error) {
	data, err := repo.dao.FindById(ctx, id)
	if err != nil {
		return domain.User{}, err
	}
	return repo.toDomain(data), nil
}

func (repo *UserCacheRepository) UpdatePassword(ctx context.Context, uid int64, oldPwd string, newPwd string) error {
	// 根据id找到对应用户 --先从缓存中寻找
	u, err := repo.cache.Get(ctx, uid)
	if err != nil {
		// 缓存中没找到数据从数据库中查询
		ur, err := repo.dao.FindById(ctx, uid)
		if err != nil {
			// 这个分支是缓存中和数据库中都没有数据
			return err
		}
		u = repo.toDomain(ur)
	}
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(oldPwd))
	if err != nil {
		return ErrInvalidPassword
	}
	// 生成新密码，并更新新密码到数据库(使用dao层面已有的update方法)
	newPassword, err := bcrypt.GenerateFromPassword([]byte(newPwd), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := domain.User{
		Id:       u.Id,
		Password: string(newPassword),
	}
	// 删除原有缓存，防止下次登录时访问到旧数据
	_ = repo.cache.Del(ctx, uid)
	return repo.dao.UpdateNonZeroFields(ctx, repo.toEntity(user))
}
func (repo *UserCacheRepository) FindById(ctx context.Context, uid int64) (domain.User, error) {
	// 先从缓存中寻找数据
	user, err := repo.cache.Get(ctx, uid)
	if err != nil {
		u, err := repo.dao.FindById(ctx, uid)
		if err != nil {
			return domain.User{}, err
		}
		user = repo.toDomain(u)
	}
	// 然后将数据放进缓存， 忽略这里的错误，存进去或没有存进去 影响不大
	_ = repo.cache.Set(ctx, user)
	return user, err
}

func (repo *UserCacheRepository) Update(ctx context.Context, user domain.User) error {
	err := repo.dao.UpdateNonZeroFields(ctx, repo.toEntity(user))
	if err != nil {
		return err
	}
	// 更新后同时更新缓存
	_ = repo.cache.Set(ctx, user)
	return err
}

func (repo *UserCacheRepository) Create(ctx context.Context, user domain.User) error {
	// 注册时先创建账号后更新个人数据，这个地方拆分开了，要求前端创建完账号直接跳转更新用户信息页面
	err := repo.dao.Insert(ctx, dao.User{
		Email:    user.Email,
		Password: user.Password,
	})
	if err != nil {
		return err
	}
	// 注册后就把数据缓存下来吗？
	//-- 按照上述步骤拆分的情况，缓存下来也只有账号密码
	// 所以不引入缓存，因为更新完用户数据之后会设置缓存
	//_ = repo.cache.Set(ctx, user)
	return err

}
func (repo *UserCacheRepository) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	u, err := repo.dao.FindByEmail(ctx, email)
	if err != nil {
		return domain.User{}, err
	}
	return repo.toDomain(u), nil
}

func (repo *UserCacheRepository) toEntity(u domain.User) dao.User {
	return dao.User{
		Id:       u.Id,
		Email:    u.Email,
		Password: u.Password,
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
		// 将毫秒时间戳转换为日期格式
		Ctime: time.UnixMilli(u.Ctime),
		Utime: time.UnixMilli(u.Utime),
	}
}
