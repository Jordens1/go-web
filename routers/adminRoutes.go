package routers

import (
	"encoding/xml"
	"net/http"

	"github.com/Jordens1/go-web/utils/model"
	"github.com/gin-gonic/gin"
)

func AdminRoutersInit(r *gin.Engine) {

	adminRouters := r.Group("/api")
	{
		// get方法,显示表单
		adminRouters.GET("/get/form", func(c *gin.Context) {
			c.HTML(http.StatusOK, "guest/index.html", gin.H{})
		})

		// post，数据字段和index.html的表单对应
		adminRouters.POST("/post/doAddUser", func(c *gin.Context) {
			username := c.PostForm("username")
			pass := c.DefaultPostForm("passwd", "woshishei")
			id := c.DefaultPostForm("id", "12")
			c.JSON(http.StatusOK, gin.H{
				"name": username,
				"pass": pass,
				"id":   id,
			})
		})

		// 获取表单的数据绑定到结构体上
		adminRouters.GET("/get/form/bind", func(c *gin.Context) {
			p := &model.People{}

			if err := c.ShouldBind(&p); err == nil {
				c.JSON(http.StatusOK, p)
			} else {
				c.JSON(http.StatusOK, gin.H{
					"err": err.Error(),
				})
			}
		})
		adminRouters.POST("/post/doAddUser/withStruct", func(c *gin.Context) {
			p := &model.People{}
			if err := c.ShouldBind(&p); err == nil {
				c.JSON(http.StatusOK, p)
			} else {
				c.JSON(http.StatusBadRequest, gin.H{
					"err": err.Error(),
				})
			}
		})

		// xml数据的绑定
		adminRouters.POST("/post/xml/bind", func(c *gin.Context) {
			art := &model.Article2{}

			b, _ := c.GetRawData()
			if err := xml.Unmarshal(b, &art); err == nil {
				c.JSON(http.StatusOK, art)
			} else {
				c.JSON(http.StatusBadRequest, gin.H{
					"err": err.Error(),
				})
			}
		})

		// json数据
		adminRouters.POST("/json", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"xishi": "name",
			})
		})

		article := &Article{
			Title:   "test",
			Desc:    "这还是测试",
			Content: "xishdic sicdsc",
		}

		// json数据，带参数
		adminRouters.GET("/json1", func(c *gin.Context) {
			c.JSON(http.StatusOK, article)
		})

		// jsonp方法，xxx/?callback=xxx
		adminRouters.GET("/jsonp", func(c *gin.Context) {
			c.JSONP(http.StatusOK, article)
		})

		// 返回xml
		adminRouters.GET("/xml", func(c *gin.Context) {
			c.XML(http.StatusOK, article)
		})

		// 获取多层的模板的，指定名字，对应模板的名字
		adminRouters.GET("/tmp", func(c *gin.Context) {

			c.HTML(http.StatusOK, "admin/index.html", gin.H{
				"title":   "我是模板",
				"another": article,
				"date":    12768463,
			})
		})

	}
}
