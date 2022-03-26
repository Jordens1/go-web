package admin

import (
	"log"
	"net/http"
	"path"

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

// 单文件上传
func (uc *UserController) UserAddUpload(c *gin.Context) {
	username := c.PostForm("username")
	file, err := c.FormFile("face")
	filename := file.Filename
	if err == nil {
		dst := path.Join("./static/upload", filename)
		c.SaveUploadedFile(file, dst)
	}

	c.JSON(http.StatusOK, gin.H{
		"success": 200,
		"name":    username,
	})

}

// 多文件上传
func (uc *UserController) UserAddUploadMutil(c *gin.Context) {
	username := c.PostForm("username")
	file, err := c.FormFile("face")
	filename := file.Filename
	if err == nil {
		dst := path.Join("./static/upload", filename)
		c.SaveUploadedFile(file, dst)
	}

	// 不同名字的文件,再写一遍单文件的上传方式
	file2, err := c.FormFile("face2")
	filename2 := file2.Filename
	if err == nil {
		dst := path.Join("./static/upload", filename2)
		c.SaveUploadedFile(file2, dst)
	}

	// 相同名字的文件上传
	form, _ := c.MultipartForm()
	files := form.File["face2"]
	abs_path := "./static/upload/mutil/"
	for _, file := range files {
		log.Println(file.Filename)
		c.SaveUploadedFile(file, path.Join(abs_path, file.Filename))
	}

	c.JSON(http.StatusOK, gin.H{
		"success": 200,
		"name":    username,
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
