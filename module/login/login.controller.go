package login

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"houserqu.com/gin-starter/core"
)

type LoginDto struct {
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// @Summary 手机号密码登录
// @Tags 登录
// @Param params body login.LoginDto true "cansh "
// @Success 200 {number} 1
// @Router /api/login/phone [post]
func LoginByPhone(c *gin.Context) {
	userId := 1

	// 参数校验
	var loginDto LoginDto
	if err := c.ShouldBindJSON(&loginDto); err != nil {
		core.ResError(c, core.ErrParam, err.Error())
		return
	}

	// 查询用户
	user, err := CheckUserByPhoneAndPassword(c, loginDto.Phone, loginDto.Password)
	if err != nil {
		core.ResError(c, core.ErrLoginFail, err.Error())
		return
	}

	// 查询成功，写登录态
	session := sessions.Default(c)
	session.Set("userId", user.ID)
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
