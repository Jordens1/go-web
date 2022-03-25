package admin

import (
	"encoding/xml"
	"net/http"

	"github.com/Jordens1/go-web/utils/model"
	"github.com/gin-gonic/gin"
)

type FormController struct{}

func (uc *FormController) GetForm(c *gin.Context) {
	c.HTML(http.StatusOK, "guest/index.html", gin.H{})
}

func (uc *FormController) GetFormBind(c *gin.Context) {
	p := &model.People{}

	if err := c.ShouldBind(&p); err == nil {
		c.JSON(http.StatusOK, p)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})
	}
}

func (uc *FormController) PostXmlBind(c *gin.Context) {
	art := &model.Article2{}

	b, _ := c.GetRawData()
	if err := xml.Unmarshal(b, &art); err == nil {
		c.JSON(http.StatusOK, art)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
	}
}

func (uc *FormController) PostJson(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"xishi": "name",
	})
}

func (uc *FormController) Json1(c *gin.Context) {
	c.JSON(http.StatusOK, model.Article{})
}

func (uc *FormController) Jsonp(c *gin.Context) {
	c.JSONP(http.StatusOK, model.Article{})
}

func (uc *FormController) Xml(c *gin.Context) {
	c.XML(http.StatusOK, model.Article{})
}

func (uc *FormController) Tmp(c *gin.Context) {

	c.HTML(http.StatusOK, "admin/index.html", gin.H{
		"title":   "我是模板",
		"another": model.Article{},
		"date":    12768463,
	})
}
