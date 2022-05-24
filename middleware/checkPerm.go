package middleware

import (
	"github.com/gin-gonic/gin"
	"houserqu.com/tiger/constants"
	"houserqu.com/tiger/core"
	"houserqu.com/tiger/module/auth"
)

// 校验当前登录用户是否有权限
func CheckPerm(needPermissions ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userPermissions, err := auth.GetUserPermissions(c.GetUint("userId")) // 登录中间件 set
		if err != nil {
			core.ResError(c, constants.ErrNoPermission, err.Error())
			c.Abort()
			return
		}

		for _, userPermission := range userPermissions {
			// 如果是超级管理员权限，直接通过
			if userPermission == "ADMIN" {
				c.Next()
				return
			}

			// 非超级管理员权限，与接口权限校验
			for _, needPermission := range needPermissions {
				if needPermission == userPermission {
					c.Next()
					return
				}
			}
		}

		core.ResError(c, constants.ErrNoPermission, "")
		c.Abort()
		return
	}
}
