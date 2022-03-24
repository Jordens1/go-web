package main

import (
	"html/template"
	"time"

	"github.com/Jordens1/go-web/routers"
	"github.com/gin-gonic/gin"
)

// 模板函数的实现
func UnixToTime(ts int) string {
	t := time.Unix(int64(ts), 0)
	return t.Format("2006-01-02 15:04:05")
}

func main() {
	r := gin.Default()

	// 静态的web库，第一个参数为路由xxx.在模板中就可使用这个地址了
	r.Static("/xxx", "./static")

	// 自定义模板函数,直接在模板中调用
	r.SetFuncMap(template.FuncMap{
		"UnixToTime": UnixToTime,
	})

	// 加载模板路径
	r.LoadHTMLGlob("templates/**/*")

	// 初始化路由
	routers.AdminRoutersInit(r)
	routers.ApiRoutersInit(r)
	routers.DefaultRoutersInit(r)

	r.Run()
}
