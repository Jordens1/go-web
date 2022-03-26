package routers

import (
	"github.com/Jordens1/go-web/controllers/api"
	"github.com/Jordens1/go-web/middleware"

	"github.com/gin-gonic/gin"
)

func ApiRoutersInit(r *gin.Engine) {
	ac := &api.ApiController{}
	apiRouters := r.Group("/api", middleware.InitMiddleware3)
	{
		// get动态路由
		apiRouters.GET("/para/:uid", ac.ParaUid)

	}

}
