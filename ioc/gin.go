package ioc

import (
	_ "github.com/StarJoice/tech_blog/docs"
	"github.com/StarJoice/tech_blog/internal/article"
	"github.com/StarJoice/tech_blog/internal/comment"
	"github.com/StarJoice/tech_blog/internal/interactive"
	"github.com/StarJoice/tech_blog/internal/label"
	"github.com/StarJoice/tech_blog/internal/search"
	"github.com/StarJoice/tech_blog/internal/user"
	"github.com/StarJoice/tools/ginx/session"
	"github.com/gin-contrib/cors"
	"github.com/gotomicro/ego/server/egin"
	"strings"
)

func InitGinXServer(sp session.Provider,
	user *user.Handler,
	arts *article.Handler,
	lab *label.Handler,
	inter *interactive.Handler,
	com *comment.Handler,
	search *search.Handler) *egin.Component {
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
	// 暂时注释掉，开发环境暂时不需要
	//server.GET("/metrics", gin.WrapH(promhttp.Handler()))
	user.PublicRoutes(server.Engine)
	arts.PublicRoutes(server.Engine)
	lab.PublicRoutes(server.Engine)
	// 开启登录校验
	server.Use(session.CheckLoginMiddleware())
	inter.PrivateRoutes(server.Engine)
	user.PrivateRoutes(server.Engine)
	arts.PrivateRoutes(server.Engine)
	com.PrivateRoutes(server.Engine)
	search.PrivateRoutes(server.Engine)

	return server
}
