package middleware

import (
	"github.com/gin-gonic/gin"
	"houserqu.com/gin-starter/core"
	"houserqu.com/gin-starter/module/auth"
)

// 校验当前登录用户是否有权限
func CheckPerm(permission string) gin.HandlerFunc {
	return func(c *gin.Context) {
		permissions, err := auth.GetUserPermissions(c.GetInt("userId")) // 登录中间件 set
		if err != nil {
			core.ResError(c, core.ErrNoPermission, err.Error())
			c.Abort()
			return
		}

		for _, v := range permissions {
			if permission == v {
				c.Next()
				return
			}
		}

		core.ResError(c, core.ErrNoPermission, "")
		c.Abort()
		return
	}
}
