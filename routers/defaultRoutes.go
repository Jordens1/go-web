package routers

import (
	"github.com/Jordens1/go-web/controllers/classic"
	"github.com/gin-gonic/gin"
)

func DefaultRoutersInit(r *gin.Engine) {
	cc := &classic.ClassicController{}
	defaultRouters := r.Group("/default")
	{
		// get方法
		defaultRouters.GET("/", cc.ParaGet1)

		// get方法,传值
		defaultRouters.GET("/para", cc.ParaGet2)

	}

}
