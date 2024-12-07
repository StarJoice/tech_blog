//@Date 2024/12/5 01:11
//@Desc

package ioc

import (
	"github.com/StarJoice/tech_blog/internal/user/web"
	"github.com/StarJoice/tools/ginx/session"
	"github.com/gin-contrib/cors"
	"github.com/gotomicro/ego/server/egin"
	"strings"
)

func InitGinXServer(sp session.Provider, user *web.UserHandler) *egin.Component {
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
	// 开启登录校验
	server.Use(session.CheckLoginMiddleware())
	return server
}
