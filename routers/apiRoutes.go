package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ApiRoutersInit(r *gin.Engine) {
	apiRouters := r.Group("/api")
	{
		// get动态路由
		apiRouters.GET("/para/:uid", func(c *gin.Context) {
			uid := c.Param("uid")

			c.JSON(http.StatusOK, gin.H{
				"uid": uid,
			})
		})

	}

}
