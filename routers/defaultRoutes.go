package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DefaultRoutersInit(r *gin.Engine) {

	defaultRouters := r.Group("/")
	{
		// get方法
		defaultRouters.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "值是：%v", "你好")
		})

		// get方法,传值
		defaultRouters.GET("/para", func(c *gin.Context) {
			name := c.Query("name")
			age := c.DefaultQuery("age", "18")

			c.JSON(http.StatusOK, gin.H{
				"name": name,
				"age":  age,
			})
		})

	}

}
