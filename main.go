package main

import (
	"html/template"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

// 模板函数的实现
func UnixToTime(ts int) string {
	t := time.Unix(int64(ts), 0)
	return t.Format("2006-01-02 15:04:05")
}

func main() {
	r := gin.Default()

	// 静态的web库，第一个参数为路由xxx.在模板中就可使用这个地址了
	r.Static("/xxx", "./static")

	// 自定义模板函数,直接在模板中调用
	r.SetFuncMap(template.FuncMap{
		"UnixToTime": UnixToTime,
	})
	// 加载模板路径
	r.LoadHTMLGlob("templates/**/*")

	// get方法
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "值是：%v", "你好")
	})

	// json数据
	r.POST("/json", func(c *gin.Context) {
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
	r.GET("/json1", func(c *gin.Context) {
		c.JSON(http.StatusOK, article)
	})

	// jsonp方法，xxx/?callback=xxx
	r.GET("/jsonp", func(c *gin.Context) {
		c.JSONP(http.StatusOK, article)
	})

	// 返回xml
	r.GET("/xml", func(c *gin.Context) {
		c.XML(http.StatusOK, article)
	})

	// 获取多层的模板的，指定名字，对应模板的名字
	r.GET("/tmp", func(c *gin.Context) {

		c.HTML(http.StatusOK, "admin/index.html", gin.H{
			"title":   "我是模板",
			"another": article,
			"date":    12768463,
		})
	})

	r.Run()
}
