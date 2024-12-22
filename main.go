package main

import (
	"github.com/StarJoice/tech_blog/ioc"
	"github.com/gotomicro/ego"
)

// go run main.go --config=config/config.yaml
//
//export EGO_DEBUG=true
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
