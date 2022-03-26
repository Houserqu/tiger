package module

import (
	"github.com/gin-gonic/gin"
	"houserqu.com/gin-starter/middleware"
	"houserqu.com/gin-starter/module/login"
	"houserqu.com/gin-starter/module/user"
	"houserqu.com/gin-starter/module/view"
)

func InitRouter(r *gin.Engine) {
	// 首页
	r.GET("/admin", view.IndexView)

	// 无需鉴权的 api接口
	r.POST("/api/login/phone", login.LoginByPhone) // 登录

	// 需要鉴权的 api接口
	api := r.Group("api", middleware.CheckLogin())
	{
		api.GET("/login/logout", login.Logout)                    // 注销
		api.GET("/login/adminloginInfo", login.GetAdminLoginInfo) // 用户信息

		// 用户相关
		api.GET("/user/:id", user.GetUser)                                        // 查单个
		api.GET("/user/list", middleware.CheckPerm("USER_ALL"), user.GetUserList) // 查列表
		api.POST("/user/create", user.CreateUser)                                 // 创建
		api.POST("/user/update", user.UpdateUser)                                 // 更新
		api.POST("/user/delete/:id", user.DeleteUser)                             // 删除
	}
}
