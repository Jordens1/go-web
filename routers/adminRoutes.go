package routers

import (
	"fmt"

	"github.com/Jordens1/go-web/controllers/admin"
	"github.com/gin-gonic/gin"
)

func AdminRoutersInit(r *gin.Engine) {

	uc := &admin.UserController{}
	fc := &admin.FormController{}

	adminRouters := r.Group("/admin")
	{

		// get方法,显示表单
		adminRouters.GET("/", func(ctx *gin.Context) {
			fmt.Println("xishi")
		}, uc.UserIndex)

		// get方法,显示表单
		adminRouters.GET("/user", uc.UserList)

		// post，数据字段和index.html的表单对应
		adminRouters.GET("/user/add", uc.UserAdd)

		// post，数据字段和index.html的表单对应,文件上传
		adminRouters.POST("/user/add/upload", uc.UserAddUpload)
		// get方法,显示表单
		adminRouters.GET("/user/edit", func(c *gin.Context) {
			c.String(200, "用户列表-edit")
		})

		// post，数据字段和index.html的表单对应,多文件上传
		adminRouters.POST("/user/add/upload/mutil", uc.UserAddUploadMutil)

		// get方法,显示表单
		adminRouters.GET("/get/form/:name", fc.GetForm)

		// 获取表单的数据绑定到结构体上
		adminRouters.GET("/get/form/bind", fc.GetFormBind)
		adminRouters.POST("/post/doAddUser/withStruct", uc.UserAddWithStruct)

		// xml数据的绑定
		adminRouters.POST("/post/xml/bind", fc.PostXmlBind)

		// json数据
		adminRouters.POST("/json", fc.PostJson)

		// json数据，带参数
		adminRouters.GET("/json1", fc.Json1)

		// jsonp方法，xxx/?callback=xxx
		adminRouters.GET("/jsonp", fc.Jsonp)

		// 返回xml
		adminRouters.GET("/xml", fc.Xml)

		// 获取多层的模板的，指定名字，对应模板的名字
		adminRouters.GET("/tmp", fc.Tmp)

	}
}
