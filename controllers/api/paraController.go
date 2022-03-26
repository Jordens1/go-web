package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiController struct{}

func (ac *ApiController) ParaUid(c *gin.Context) {
	uid := c.Param("uid")

	// 获取中间件里的共享数据
	conText, exits := c.Get("con")
	if exits {
		c.JSON(http.StatusOK, gin.H{
			"uid":     uid,
			"conText": conText,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"uid": uid,
		})
	}

}
