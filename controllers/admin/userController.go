package admin

import (
	"net/http"

	"github.com/Jordens1/go-web/utils/model"
	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (uc *UserController) UserIndex(c *gin.Context) {
	c.String(200, "后台首页")
}

func (uc *UserController) UserList(c *gin.Context) {
	c.String(200, "用户列表")
}

func (uc *UserController) UserAdd(c *gin.Context) {
	username := c.PostForm("username")
	pass := c.DefaultPostForm("passwd", "woshishei")
	id := c.DefaultPostForm("id", "12")
	c.JSON(http.StatusOK, gin.H{
		"name": username,
		"pass": pass,
		"id":   id,
	})
}

func (uc *UserController) GetForm(c *gin.Context) {
	c.HTML(http.StatusOK, "guest/index.html", gin.H{})
}

func (uc *UserController) UserAddWithStruct(c *gin.Context) {
	p := &model.People{}
	if err := c.ShouldBind(&p); err == nil {
		c.JSON(http.StatusOK, p)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
	}
}
