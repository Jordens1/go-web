package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiController struct{}

func (ac *ApiController) ParaUid(c *gin.Context) {
	uid := c.Param("uid")

	c.JSON(http.StatusOK, gin.H{
		"uid": uid,
	})
}
