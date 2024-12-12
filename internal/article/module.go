//@Date 2024/12/12 16:25
//@Desc

package article

import (
	"github.com/StarJoice/tech_blog/internal/article/service"
	"github.com/StarJoice/tech_blog/internal/article/web"
)

type Service = service.Service
type Handler = web.ArticleHandler

type Module struct {
	Svc Service
	Hdl *Handler
}
