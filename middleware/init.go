package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// 中间件
func InitMiddleware(ctx *gin.Context) {
	fmt.Println("中件")
	fmt.Println(ctx.Request.URL)
	ctx.Next()
	// ctx.Abort()
	fmt.Println("next")
}

func InitMiddleware2(ctx *gin.Context) {
	fmt.Println("中件2")
	// ctx.Next()
	ctx.Abort()
	fmt.Println("next2")
}

// 演示通过上下文来共享数据
func InitMiddleware3(ctx *gin.Context) {
	fmt.Println("中件3")
	// ctx.Next()
	ctx.Set("con", "woshi3")

	cCp := ctx.Copy()
	// 在另一个goroutine中,只能使用上下文的只读副本.不然的话,闭包中的数据会变化
	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println("Done in :" + cCp.Request.RequestURI)
	}()
}
