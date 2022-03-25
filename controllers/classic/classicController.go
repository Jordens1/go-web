package classic

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ClassicController struct{}

func (ac *ClassicController) ParaGet1(c *gin.Context) {
	c.String(http.StatusOK, "值是：%v", "你好")
}

func (ac *ClassicController) ParaGet2(c *gin.Context) {
	name := c.Query("name")
	age := c.DefaultQuery("age", "18")

	c.JSON(http.StatusOK, gin.H{
		"name": name,
		"age":  age,
	})
}
