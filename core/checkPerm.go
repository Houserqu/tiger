package core

import (
	"github.com/gin-gonic/gin"
	"houserqu.com/gin-starter/module/perm"
)

// 校验当前登录用户是否有权限
func CheckPerm(permission string) gin.HandlerFunc {
	return func(c *gin.Context) {
		permissions, err := perm.GetUserPermissions(c.GetInt("userId")) // 登录中间件 set
		if err != nil {
			ResError(c, ErrNoPermission, err.Error())
		}

		for _, v := range permissions {
			if permission == v {
				c.Next()
				return
			}
		}

		ResError(c, ErrNoPermission, "")
		c.Abort()
		return
	}
}
