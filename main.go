package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/Jordens1/go-web/middleware"
	"github.com/Jordens1/go-web/routers"
	"github.com/gin-gonic/gin"
)

// 模板函数的实现
func UnixToTime(ts int) string {
	t := time.Unix(int64(ts), 0)
	return t.Format("2006-01-02 15:04:05")
}

func main() {
	// 该方法已经含有了两个默认的中间件,不想使用的话,可以用 gin.New()
	r := gin.Default()

	// 静态的web库，第一个参数为路由xxx.在模板中就可使用这个地址了
	r.Static("/xxx", "./static")

	// 自定义模板函数,直接在模板中调用
	r.SetFuncMap(template.FuncMap{
		"UnixToTime": UnixToTime,
	})

	// 加载模板路径
	r.LoadHTMLGlob("templates/**/*")

	// 全局中间件
	// r.Use(InitMiddleware, InitMiddleware2)
	// 增减中间件
	r.GET("/xishi", middleware.InitMiddleware2, func(ctx *gin.Context) {
		fmt.Println("handler")
		ctx.String(http.StatusOK, "200")
	})

	// 初始化路由
	routers.AdminRoutersInit(r)
	routers.ApiRoutersInit(r)
	routers.DefaultRoutersInit(r)

	r.Run()
}
