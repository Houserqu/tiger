package login

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"houserqu.com/gin-starter/core"
)

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
