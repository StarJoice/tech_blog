package service

import (
	"context"
	"errors"
	"github.com/StarJoice/tech_blog/internal/user/internal/domain"
	"github.com/StarJoice/tech_blog/internal/user/internal/repository"
	repomocks "github.com/StarJoice/tech_blog/internal/user/internal/repository/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestUserSvc_Signup(t *testing.T) {
	TestCases := []struct {
		name string
		mock func(ctrl *gomock.Controller) repository.UserRepository

		ctx     context.Context
		user    domain.User
		WantErr error
	}{
		{
			name: "注册成功",
			mock: func(ctrl *gomock.Controller) repository.UserRepository {
				repo := repomocks.NewMockUserRepository(ctrl)
				repo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil)
				return repo
			},
			ctx: context.Background(),
			user: domain.User{
				Email:    "123456@test.com",
				Password: "test@123",
			},
			WantErr: nil,
		},
		{
			name: "注册失败--用户已存在",
			mock: func(ctrl *gomock.Controller) repository.UserRepository {
				repo := repomocks.NewMockUserRepository(ctrl)
				repo.EXPECT().Create(gomock.Any(), gomock.Any()).
					Return(repository.ErrDuplicateEmail)
				return repo
			},
			ctx: context.Background(),
			user: domain.User{
				Email:    "123456@test.com",
				Password: "test@123",
			},
			WantErr: repository.ErrDuplicateEmail,
		},
		{
			name: "注册失败--未知错误",
			mock: func(ctrl *gomock.Controller) repository.UserRepository {
				repo := repomocks.NewMockUserRepository(ctrl)
				repo.EXPECT().Create(gomock.Any(), gomock.Any()).
					Return(errors.New("test error"))
				return repo
			},
			ctx: context.Background(),
			user: domain.User{
				Email:    "123456@test.com",
				Password: "test@123",
			},
			WantErr: errors.New("test error"),
		},
	}
	for _, tc := range TestCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := tc.mock(ctrl)
			svc := NewUserSvc(repo)
			err := svc.Signup(tc.ctx, tc.user)
			assert.Equal(t, tc.WantErr, err)
		})
	}
}

func TestUserSvc_Login(t *testing.T) {
	TestCases := []struct {
		name     string
		mock     func(ctrl *gomock.Controller) repository.UserRepository
		ctx      context.Context
		email    string
		password string
		wantErr  error
		WantRes  domain.User
	}{
		{
			name: "登录成功",
			mock: func(ctrl *gomock.Controller) repository.UserRepository {
				repo := repomocks.NewMockUserRepository(ctrl)
				repo.EXPECT().FindByEmail(gomock.Any(), gomock.Any()).
					Return(domain.User{
						Email:    "123456@test.com",
						Password: "$2a$10$u.3fTmeLVdcHBmaxyjHMyuauH32A7QzT.1B5iPc.Idcg4GdpxKCaO",
					}, nil)
				return repo
			},
			ctx:      context.Background(),
			email:    "123456@test.com",
			password: "test@123",
			WantRes: domain.User{
				Email:    "123456@test.com",
				Password: "$2a$10$u.3fTmeLVdcHBmaxyjHMyuauH32A7QzT.1B5iPc.Idcg4GdpxKCaO",
			},
			wantErr: nil,
		},
		{
			name: "登录失败--用户不存在",
			mock: func(ctrl *gomock.Controller) repository.UserRepository {
				repo := repomocks.NewMockUserRepository(ctrl)
				repo.EXPECT().FindByEmail(gomock.Any(), gomock.Any()).
					Return(domain.User{}, repository.ErrUserNotFound)
				return repo
			},
			ctx:      context.Background(),
			email:    "123456@test.com",
			password: "test@123",
			WantRes:  domain.User{},
			wantErr:  ErrInvalidUserOrPassword,
		},
		{
			name: "登录失败--密码错误",
			mock: func(ctrl *gomock.Controller) repository.UserRepository {
				repo := repomocks.NewMockUserRepository(ctrl)
				repo.EXPECT().FindByEmail(gomock.Any(), gomock.Any()).
					Return(domain.User{
						Email:    "123456@test.com",
						Password: "$2a$10$u.3fTmeLVdcHBmaxyjHMyuauH32A7QzT.1B5iPc.Idcg4GdpxKCaO",
					}, nil)
				return repo
			},
			ctx:      context.Background(),
			email:    "123456@test.com",
			password: "test@1234",
			WantRes:  domain.User{},
			wantErr:  ErrInvalidUserOrPassword,
		},
	}
	for _, tc := range TestCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			repo := tc.mock(ctrl)
			svc := NewUserSvc(repo)
			res, err := svc.Login(tc.ctx, tc.email, tc.password)
			assert.Equal(t, tc.wantErr, err)
			assert.Equal(t, tc.WantRes, res)
		})
	}
}

func TestUserSvc_Profile(t *testing.T) {
	TestCases := []struct {
		name    string
		mock    func(ctrl *gomock.Controller) repository.UserRepository
		ctx     context.Context
		uid     int64
		wantRes domain.User
		wantErr error
	}{
		{
			name: "查询成功",
			mock: func(ctrl *gomock.Controller) repository.UserRepository {
				repo := repomocks.NewMockUserRepository(ctrl)
				repo.EXPECT().FindById(gomock.Any(), gomock.Any()).
					Return(domain.User{
						Id: 1,
					}, nil)
				return repo
			},
			ctx: context.Background(),
			uid: 1,
			wantRes: domain.User{
				Id: 1,
			},
			wantErr: nil,
		},
		{
			name: "查询失败",
			mock: func(ctrl *gomock.Controller) repository.UserRepository {
				repo := repomocks.NewMockUserRepository(ctrl)
				repo.EXPECT().FindById(gomock.Any(), gomock.Any()).
					Return(domain.User{}, errors.New("test error"))
				return repo
			},
			ctx:     context.Background(),
			uid:     1,
			wantRes: domain.User{},
			wantErr: errors.New("test error"),
		},
	}
	for _, tc := range TestCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			repo := tc.mock(ctrl)
			svc := NewUserSvc(repo)
			res, err := svc.Profile(tc.ctx, tc.uid)
			assert.Equal(t, tc.wantRes, res)
			assert.Equal(t, tc.wantErr, err)
		})
	}
}

func TestUserSvc_UpdateNonSensitiveInfo(t *testing.T) {
	TestCases := []struct {
		name    string
		mock    func(ctrl *gomock.Controller) repository.UserRepository
		ctx     context.Context
		u       domain.User
		WantErr error
	}{
		{
			name: "更新成功",
			mock: func(ctrl *gomock.Controller) repository.UserRepository {
				repo := repomocks.NewMockUserRepository(ctrl)
				repo.EXPECT().Update(gomock.Any(), gomock.Any()).
					Return(nil)
				return repo
			},
			ctx: context.Background(),
			u: domain.User{
				Email:    "123456@test.com",
				Password: "test@123",
			},
			WantErr: nil,
		},
	}
	for _, tc := range TestCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			repo := tc.mock(ctrl)
			svc := NewUserSvc(repo)
			err := svc.UpdateNonSensitiveInfo(tc.ctx, tc.u)
			assert.Equal(t, tc.WantErr, err)
		})
	}
}

func TestUserSvc_UpdatePassword(t *testing.T) {
	TestCases := []struct {
		name    string
		mock    func(ctrl *gomock.Controller) repository.UserRepository
		ctx     context.Context
		uid     int64
		oldPwd  string
		newPwd  string
		wantErr error
	}{
		{
			name: "更新成功",
			mock: func(ctrl *gomock.Controller) repository.UserRepository {
				repo := repomocks.NewMockUserRepository(ctrl)
				repo.EXPECT().
					UpdatePassword(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Return(nil)
				return repo
			},
			ctx:     context.Background(),
			uid:     1,
			oldPwd:  "test1",
			newPwd:  "test2",
			wantErr: nil,
		},
	}
	for _, tc := range TestCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			repo := tc.mock(ctrl)
			svc := NewUserSvc(repo)
			err := svc.UpdatePassword(tc.ctx, tc.uid, tc.oldPwd, tc.newPwd)
			assert.Equal(t, tc.wantErr, err)
		})
	}
}

func TestUserSvc_GetByID(t *testing.T) {
	TestCases := []struct {
		name    string
		mock    func(ctrl *gomock.Controller) repository.UserRepository
		ctx     context.Context
		uid     int64
		wantRes domain.User
		wantErr error
	}{
		{
			name: "查询成功",
			mock: func(ctrl *gomock.Controller) repository.UserRepository {
				repo := repomocks.NewMockUserRepository(ctrl)
				repo.EXPECT().GetByID(gomock.Any(), gomock.Any()).
					Return(domain.User{
						Id: 1,
					}, nil)
				return repo
			},
			ctx: context.Background(),
			uid: 1,
			wantRes: domain.User{
				Id: 1,
			},
			wantErr: nil,
		},
	}
	for _, tc := range TestCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			repo := tc.mock(ctrl)
			svc := NewUserSvc(repo)
			id, err := svc.GetByID(tc.ctx, tc.uid)
			assert.Equal(t, tc.wantErr, err)
			assert.Equal(t, tc.wantRes, id)
		})
	}
}

// 生成一个加密后的用户密码，以供测试使用
func TestBcrypt(t *testing.T) {
	num := "test@123"
	res, err := bcrypt.GenerateFromPassword([]byte(num), bcrypt.DefaultCost)
	assert.NoError(t, err)
	t.Log(string(res))
}
