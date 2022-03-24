package main

import (
	"encoding/xml"
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

// get的参数会根据表单的值进行对应
type People struct {
	Name string `json:"name" form:"name"`
	Age  string `json:"age" form:"age"`
	Sex  string `json:"sex" form:"sex"`
}

// xml数据进行解析绑定
type Article2 struct {
	Title   string `json:"title" xml:"title"`
	Content string `json:"content" xml:"content"`
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

	// get方法,传值
	r.GET("/para", func(c *gin.Context) {
		name := c.Query("name")
		age := c.DefaultQuery("age", "18")

		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})

	// get动态路由
	r.GET("/para/:uid", func(c *gin.Context) {
		uid := c.Param("uid")

		c.JSON(http.StatusOK, gin.H{
			"uid": uid,
		})
	})

	// get方法,显示表单
	r.GET("/get/form", func(c *gin.Context) {
		c.HTML(http.StatusOK, "guest/index.html", gin.H{})

	})
	// post，数据字段和index.html的表单对应
	r.POST("/post/doAddUser", func(c *gin.Context) {
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
	r.GET("/get/form/bind", func(c *gin.Context) {
		p := &People{}
		if err := c.ShouldBind(&p); err == nil {
			c.JSON(http.StatusOK, p)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"err": err.Error(),
			})
		}
	})
	r.POST("/post/doAddUser/withStruct", func(c *gin.Context) {
		p := &People{}
		if err := c.ShouldBind(&p); err == nil {
			c.JSON(http.StatusOK, p)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
		}
	})

	// xml数据的绑定
	r.POST("/post/xml/bind", func(c *gin.Context) {
		art := &Article2{}

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
