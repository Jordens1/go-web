package api

import (
	"net/http"

	"github.com/Jordens1/go-web/utils/bingfa"
	"github.com/gin-contrib/sessions"
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

// 演示使用session
func (ac *ApiController) GetSession1(c *gin.Context) {
	// 设置session,并且保存
	session := sessions.Default(c)
	session.Set("sessionName", "cookie1")
	// 设置过期时间
	session.Options(sessions.Options{MaxAge: 3600 * 6})
	session.Save()

	c.String(200, "yes")

}

func (ac *ApiController) GetSession2(c *gin.Context) {

	// session获取
	session := sessions.Default(c)
	sesssionName := session.Get("sessionName")
	// c.JSON(200, gin.H{
	// 	"sessionName": sesssionName,
	// })
	c.String(200, "woshi :%v", sesssionName)
}

func (ac *ApiController) Test1(c *gin.Context) {
	bingfa.BufferChannel()
	c.String(200, "ok")
}
