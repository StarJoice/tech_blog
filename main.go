//@Date 2024/12/5 00:42
//@Desc

package main

import (
	"github.com/StarJoice/tech_blog/ioc"
	"github.com/gotomicro/ego"
)

// export EGO_DEBUG=true
// go run main.go --config=config/config.yaml
func main() {
	// 先触发初始化
	egoApp := ego.New()
	app, err := ioc.InitApp()
	if err != nil {
		panic(err)
	}
	err = egoApp.
		// Invoker 在 Ego 里面，应该叫做初始化函数
		Invoker().
		Serve(app.Web).
		Run()
	panic(err)
}
