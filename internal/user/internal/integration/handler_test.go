//go:build e2eS

package integration

import (
	"github.com/StarJoice/tech_blog/internal/test"
	testioc "github.com/StarJoice/tech_blog/internal/test/ioc"
	"github.com/StarJoice/tech_blog/internal/user/internal/integration/startup"
	"github.com/StarJoice/tech_blog/internal/user/internal/repository/dao"
	"github.com/StarJoice/tools/ginx/session"
	"github.com/ego-component/egorm"
	"github.com/gin-gonic/gin"
	"github.com/gotomicro/ego/core/econf"
	"github.com/gotomicro/ego/server/egin"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type HandlerTestSuite struct {
	suite.Suite
	db     *egorm.Component
	server *egin.Component
}

func (s *HandlerTestSuite) SetupSuite() {
	econf.Set("http_users", map[string]any{})
	s.db = testioc.InitDB()
	err := dao.InitTable(s.db)
	require.NoError(s.T(), err)
	econf.Set("server", map[string]string{})
	server := egin.Load("server").Build()
	hdl := startup.InitHandler(nil, nil)
	server.Use(func(ctx *gin.Context) {
		ctx.Set("_session", session.NewMemorySession(session.Claims{
			Uid: 123,
		}))
	})
	hdl.PrivateRoutes(server.Engine)
	s.server = server
}
func (s *HandlerTestSuite) TearDownSuite() {
	err := s.db.Exec("TRUNCATE table `user`").Error
	require.NoError(s.T(), err)
}

func (s *HandlerTestSuite) TestUserHandler_Edit() {
	TestCases := []struct {
		name     string
		before   func(t *testing.T)
		after    func(t *testing.T)
		req      editReq
		wantResp test.Result[any]
		wantErr  error
	}{
		{},
	}
	for _, tc := range TestCases {
		s.T().Run(tc.name, func(t *testing.T) {
			// 准备数据
			tc.before(t)

			// 校验
			tc.after(t)
			// 清理数据

		})
	}
}

func TestUserHandler_Login(t *testing.T) {

}

func TestUserHandler_Profile(t *testing.T) {

}

func TestUserHandler_SignUp(t *testing.T) {

}

func TestUserHandler_UpdatePassword(t *testing.T) {

}
