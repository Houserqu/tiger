package middleware

import (
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"houserqu.com/gin-starter/core"
)

func CheckLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 非登录请求，校验登录态
		if c.FullPath() != "/login/login" {
			session := sessions.Default(c)
			userId := session.Get("userId")
			fmt.Println(userId)
			if userId == nil {
				core.ResError(c, core.ErrNotLogin, "")
				c.Abort()
				return
			}

			// 可以通过上下文对象，设置一些依附在上下文对象里面的键/值数据
			c.Set("userId", userId)
		}

		// 调用下一个中间件，或者控制器处理函数，具体得看注册了多少个中间件。
		c.Next()
	}
}
