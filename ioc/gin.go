//@Date 2024/12/5 01:11
//@Desc

package ioc

import (
	_ "github.com/StarJoice/tech_blog/docs"
	"github.com/StarJoice/tech_blog/internal/article"
	"github.com/StarJoice/tech_blog/internal/user"
	"github.com/StarJoice/tools/ginx/session"
	"github.com/gin-contrib/cors"
	"github.com/gotomicro/ego/server/egin"
	"strings"
)

func InitGinXServer(sp session.Provider,
	user *user.Handler, arts *article.Handler) *egin.Component {
	session.SetDefaultProvider(sp)
	server := egin.Load("Web").Build()
	//server.Use(mdls...)
	// 设置跨域
	server.Use(cors.New(cors.Config{
		ExposeHeaders:    []string{"X-Refresh-Token", "X-Access-Token"},
		AllowCredentials: true,
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		AllowOriginFunc: func(origin string) bool {
			if strings.HasPrefix(origin, "http://localhost") {
				return true
			}
			// 只允许我的域名过来的
			//return strings.Contains(origin, "")
			return false
		},
	}))
	user.PublicRoutes(server.Engine)
	arts.PublicRoutes(server.Engine)
	// 开启登录校验
	server.Use(session.CheckLoginMiddleware())
	user.PrivateRoutes(server.Engine)
	arts.PrivateRoutes(server.Engine)

	return server
}
