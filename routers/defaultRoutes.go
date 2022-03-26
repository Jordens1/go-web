package routers

import (
	"github.com/Jordens1/go-web/controllers/classic"
	"github.com/gin-gonic/gin"
)

func DefaultRoutersInit(r *gin.Engine) {

	cc := &classic.ClassicController{}

	// 配置中间件的两种写法
	// defaultRouters := r.Group("/default", middleware.InitMiddleware)
	defaultRouters := r.Group("/default")
	// defaultRouters.Use(middleware.InitMiddleware2)

	{
		// get方法
		defaultRouters.GET("/", cc.ParaGet1)

		// get方法,传值
		defaultRouters.GET("/para", cc.ParaGet2)

	}

}
