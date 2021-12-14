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
	r.GET("/", view.IndexView)

	r.GET("/example", example.GetModel)
	r.GET("/error", example.ErrorExample)

	// 登录相关
	loginRouter := r.Group("login")
	loginRouter.POST("/login", login.Login)        // 登录
	loginRouter.POST("/logout", login.Logout)      // 注销
	loginRouter.GET("/loginInfo", login.LoginInfo) // 用户信息

	// 用户相关
	userRouter := r.Group("user")
	userRouter.GET("/:id", user.GetUser)                                        // 查单个
	userRouter.GET("/list", middleware.CheckPerm("USER_ALL"), user.GetUserList) // 查列表
	userRouter.POST("/create", user.CreateUser)                                 // 创建
	userRouter.POST("/update", user.UpdateUser)                                 // 更新
	userRouter.POST("/delete/:id", user.DeleteUser)                             // 删除
}
