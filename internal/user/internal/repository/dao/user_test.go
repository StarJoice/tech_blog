package dao

import (
	"context"
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	gormMysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestUserGormDao_FindByEmail(t *testing.T) {
	testCases := []struct {
		name     string
		mock     func(t *testing.T) *sql.DB
		ctx      context.Context
		email    string
		wantUser User
		wantErr  error
	}{
		{
			name: "查询成功",
			mock: func(t *testing.T) *sql.DB {
				mockDB, mock, err := sqlmock.New()
				require.NoError(t, err)
				rows := sqlmock.NewRows([]string{"id", "email", "password", "nickname", "avatar", "about_me", "ctime", "utime"})
				rows.AddRow(1, "123456@test.com", "123", "test", "1", "1", 0, 0)
				mock.ExpectQuery("^SELECT \\* FROM `user` WHERE email = \\? ORDER BY `user`.`id` LIMIT \\?").
					WillReturnRows(rows)
				return mockDB
			},
			ctx:   context.Background(),
			email: "123456@test.com",
			wantUser: User{
				Id:       1,
				Email:    "123456@test.com",
				Password: "123",
				Nickname: "test",
				Avatar:   "1",
				AboutMe:  "1",
				Ctime:    0,
				Utime:    0,
			},
			wantErr: nil,
		},
		{
			name: "查询失败--邮箱不存在",
			mock: func(t *testing.T) *sql.DB {
				mockDB, mock, err := sqlmock.New()
				require.NoError(t, err)
				rows := sqlmock.NewRows([]string{"id", "email", "password", "nickname", "avatar", "about_me", "ctime", "utime"})
				//rows.AddRow(1, "123456@test.com", "123", "test", "1", "1", 0, 0)
				mock.ExpectQuery("^SELECT \\* FROM `user` WHERE email = \\? ORDER BY `user`.`id` LIMIT \\?").
					WillReturnRows(rows)
				return mockDB
			},
			ctx:      context.Background(),
			email:    "123456@test.com",
			wantUser: User{},
			wantErr:  gorm.ErrRecordNotFound,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db, err := gorm.Open(gormMysql.New(gormMysql.Config{
				Conn: tc.mock(t),
				// 如果为 false ，则GORM在初始化时，会先调用 show version
				SkipInitializeWithVersion: true,
			}), &gorm.Config{
				// 如果为 true ，则不允许 Ping数据库
				DisableAutomaticPing: true,
				// 如果为 false ，则即使是单一语句，也会开启事务
				SkipDefaultTransaction: true,
			})
			require.NoError(t, err)
			dao := NewUserGormDao(db)
			u, err := dao.FindByEmail(tc.ctx, tc.email)
			assert.Equal(t, tc.wantErr, err)
			assert.Equal(t, tc.wantUser, u)
		})
	}
}

func TestUserGormDao_FindById(t *testing.T) {
	testCases := []struct {
		name     string
		ctx      context.Context
		id       int64
		mock     func(t *testing.T) *sql.DB
		wantErr  error
		wantUser User
	}{
		{
			name: "查找成功",
			ctx:  context.Background(),
			id:   1,
			mock: func(t *testing.T) *sql.DB {
				mockDB, mock, err := sqlmock.New()
				require.NoError(t, err)
				rows := sqlmock.NewRows([]string{"id", "email", "password", "nickname", "avatar", "about_me", "ctime", "utime"})
				rows.AddRow(1, "123456@test.com", "123", "test", "1", "1", 0, 0)
				mock.ExpectQuery("^SELECT \\* FROM `user` WHERE Id=\\? ORDER BY `user`.`id` LIMIT \\?").
					WillReturnRows(rows)
				return mockDB
			},
			wantErr: nil,
			wantUser: User{
				Id:       1,
				Email:    "123456@test.com",
				Password: "123",
				Nickname: "test",
				Avatar:   "1",
				AboutMe:  "1",
				Ctime:    0,
				Utime:    0,
			},
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			db, err := gorm.Open(gormMysql.New(gormMysql.Config{
				Conn: tt.mock(t),
				// 如果为 false ，则GORM在初始化时，会先调用 show version
				SkipInitializeWithVersion: true,
			}), &gorm.Config{
				// 如果为 true ，则不允许 Ping数据库
				DisableAutomaticPing: true,
				// 如果为 false ，则即使是单一语句，也会开启事务
				SkipDefaultTransaction: true,
			})
			require.NoError(t, err)
			dao := NewUserGormDao(db)
			u, err := dao.FindById(tt.ctx, tt.id)
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.wantUser, u)
		})
	}

}

func TestUserGormDao_Insert(t *testing.T) {
	TestCases := []struct {
		name    string
		mock    func(t *testing.T) *sql.DB
		ctx     context.Context
		u       User
		wantErr error
	}{
		{
			name: "插入成功",
			mock: func(t *testing.T) *sql.DB {
				mockDB, mock, err := sqlmock.New()
				res := sqlmock.NewResult(3, 1)
				mock.ExpectExec("INSERT INTO `user` .*").
					WillReturnResult(res)
				require.NoError(t, err)
				return mockDB
			},
			ctx: context.Background(),
			u: User{
				Email: "123456@test.com",
			},
			wantErr: nil,
		},
		{
			name: "插入失败--邮箱冲突",
			mock: func(t *testing.T) *sql.DB {
				mockDB, mock, err := sqlmock.New()
				mock.ExpectExec("INSERT INTO `user` .*").
					WillReturnError(&mysql.MySQLError{
						Number: 1062,
					})
				require.NoError(t, err)
				return mockDB
			},
			ctx: context.Background(),
			u: User{
				Email: "123456@test.com",
			},
			wantErr: ErrDuplicateEmail,
		},
		{
			name: "插入失败--其他错误",
			mock: func(t *testing.T) *sql.DB {
				mockDB, mock, err := sqlmock.New()
				mock.ExpectExec("INSERT INTO `user` .*").
					WillReturnError(errors.New("test error"))
				require.NoError(t, err)
				return mockDB
			},
			ctx: context.Background(),
			u: User{
				Email: "123456@test.com",
			},
			wantErr: errors.New("test error"),
		},
	}
	for _, tc := range TestCases {
		t.Run(tc.name, func(t *testing.T) {
			db, err := gorm.Open(gormMysql.New(gormMysql.Config{
				Conn: tc.mock(t),
				// SELECT VERSION;
				SkipInitializeWithVersion: true,
			}), &gorm.Config{
				// 你 mock DB 不需要 ping
				DisableAutomaticPing: true,
				// 禁用默认事务
				SkipDefaultTransaction: true,
			})
			require.NoError(t, err)
			dao := NewUserGormDao(db)
			err = dao.Insert(tc.ctx, tc.u)
			assert.Equal(t, err, tc.wantErr)
		})
	}
}

func TestUserGormDao_UpdateNonZeroFields(t *testing.T) {
	testCases := []struct {
		name    string
		ctx     context.Context
		User    User
		mock    func(t *testing.T) *sql.DB
		wantErr error
	}{
		{
			name: "更新成功",
			mock: func(t *testing.T) *sql.DB {
				mockDB, mock, err := sqlmock.New()
				require.NoError(t, err)
				mock.ExpectExec("UPDATE `user` .*").
					WillReturnResult(sqlmock.NewResult(1, 1))
				return mockDB
			},
			ctx: context.Background(),
			User: User{
				Id:    1,
				Email: "123456@test.com",
			},
			wantErr: nil,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			db, err := gorm.Open(gormMysql.New(gormMysql.Config{
				Conn: tt.mock(t),
				// 如果为 false ，则GORM在初始化时，会先调用 show version
				SkipInitializeWithVersion: true,
			}), &gorm.Config{
				// 如果为 true ，则不允许 Ping数据库
				DisableAutomaticPing: true,
				// 如果为 false ，则即使是单一语句，也会开启事务
				SkipDefaultTransaction: true,
			})
			require.NoError(t, err)
			dao := NewUserGormDao(db)
			err = dao.UpdateNonZeroFields(tt.ctx, tt.User)
			assert.Equal(t, tt.wantErr, err)

		})
	}
}
