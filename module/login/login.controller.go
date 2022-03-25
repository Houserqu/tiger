package login

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"houserqu.com/gin-starter/core"
)

// @Summary 账号密码登录
// @Tags 登录
// @Success 200 {number} 1
// @Router /api/login [get]
func Login(c *gin.Context) {
	userId := 1

	session := sessions.Default(c)
	session.Set("userId", userId)
	session.Save()

	core.ResSuccess(c, userId)
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	core.ResSuccess(c, "success")
}

func LoginInfo(c *gin.Context) {
	session := sessions.Default(c)
	userId := session.Get("userId")

	core.ResSuccess(c, userId)
}
