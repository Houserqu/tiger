package login

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"houserqu.com/gin-starter/core"
	"houserqu.com/gin-starter/middleware"
)

type LoginDto struct {
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Controller(r *gin.Engine) {
	// 创建 group 并绑定中间件
	api := r.Group("/api/login")

	api.POST("/phone", LoginByPhone)                                       // 登录
	api.GET("/logout", middleware.CheckLogin(), Logout)                    // 注销
	api.GET("/adminLoginInfo", middleware.CheckLogin(), GetAdminLoginInfo) // 用户信息
}

// @Summary 手机号密码登录
// @Tags 登录
// @Param params body login.LoginDto true "cansh "
// @Success 200 {object} login.AdminLoginInfo
// @Router /api/login/phone [post]
func LoginByPhone(c *gin.Context) {
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

	// 获取登录信息
	var adminLoginInfo AdminLoginInfo
	err = GetAdminLoginInfoByUserId(c, &adminLoginInfo, user.ID)
	if err != nil {
		core.ResError(c, core.ErrLoginInfoFail, err.Error())
		return
	}

	// 查询成功，写登录态
	session := sessions.Default(c)
	session.Set("userId", user.ID)
	session.Save()

	core.ResSuccess(c, adminLoginInfo)
}

// @Summary 获取管理员登录信息
// @Tags 登录
// @Router /api/login/adminloginInfo [get]
// @Success 200 {object} login.AdminLoginInfo
func GetAdminLoginInfo(c *gin.Context) {
	userId := c.GetUint("userId")

	var adminLoginInfo AdminLoginInfo
	err := GetAdminLoginInfoByUserId(c, &adminLoginInfo, userId)
	if err != nil {
		core.ResError(c, core.ErrLoginInfoFail, err.Error())
		return
	}

	core.ResSuccess(c, adminLoginInfo)
}

// @Summary 注销
// @Tags 登录
// @Router /api/login/logout [get]
// @Success 200 {string} success
func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	core.ResSuccess(c, "success")
}
