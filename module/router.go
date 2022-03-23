package module

import (
	"github.com/gin-gonic/gin"
	"houserqu.com/gin-starter/middleware"
	"houserqu.com/gin-starter/module/example"
	"houserqu.com/gin-starter/module/login"
	"houserqu.com/gin-starter/module/user"
	"houserqu.com/gin-starter/module/view"
)

func InitRouter(r *gin.Engine) {
	// 首页
	r.GET("/admin", view.IndexView)

	// 无需鉴权的 api接口
	pub := r.Group("pub")
	{
		pub.POST("/login", login.Login) // 登录
		pub.GET("/example", example.GetModel)
		pub.GET("/error", example.ErrorExample)
	}

	// 需要鉴权的 api接口
	api := r.Group("api", middleware.CheckLogin())
	{
		api.POST("/logout", login.Logout)      // 注销
		api.GET("/loginInfo", login.LoginInfo) // 用户信息

		// 用户相关
		api.GET("/user/:id", user.GetUser)                                        // 查单个
		api.GET("/user/list", middleware.CheckPerm("USER_ALL"), user.GetUserList) // 查列表
		api.POST("/user/create", user.CreateUser)                                 // 创建
		api.POST("/user/update", user.UpdateUser)                                 // 更新
		api.POST("/user/delete/:id", user.DeleteUser)                             // 删除
	}
}
