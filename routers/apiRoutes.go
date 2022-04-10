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
		apiRouters.GET("/para/uid/:uid", ac.ParaUid)
		apiRouters.GET("/para/session1", ac.GetSession1)
		apiRouters.GET("/para/session2", ac.GetSession2)

		apiRouters.GET("/go/routine1", ac.Test1)
	}

}
