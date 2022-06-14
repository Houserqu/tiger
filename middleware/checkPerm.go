package middleware

import (
	"github.com/gin-gonic/gin"
	"houserqu.com/tiger/constants/ERR"
	"houserqu.com/tiger/core"
	"houserqu.com/tiger/model"
)

func getUserPermissions(userId uint) (allowPermIds []string, err error) {
	// 查用户角色
	var relRoles []model.RelUserRole
	err = core.Mysql.Where("user_id = ?", userId).Find(&relRoles).Error
	if err != nil {
		return
	}

	if len(relRoles) == 0 {
		return
	}

	var roleIds []uint
	for _, v := range relRoles {
		roleIds = append(roleIds, v.RoleID)
	}

	// 查角色权限ID
	var relRolePermissions []model.RelRolePermission
	err = core.Mysql.Where("role_id IN ?", roleIds).Find(&relRolePermissions).Error
	if err != nil {
		return
	}

	var permissionIds []string
	for _, v := range relRolePermissions {
		permissionIds = append(permissionIds, v.PermissionID)
	}

	// 查权限列表
	var permissions []model.Permission
	err = core.Mysql.Where("id IN ?", permissionIds).Find(&permissions).Error
	if err != nil {
		return
	}

	for _, v := range permissions {
		allowPermIds = append(allowPermIds, v.ID)
	}
	return
}

// 校验当前登录用户是否有权限
func CheckPerm(needPermissions ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userPermissions, err := getUserPermissions(c.GetUint("userId")) // 登录中间件 set
		if err != nil {
			core.ResError(c, ERR.NoPermission, err.Error())
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

		core.ResError(c, ERR.NoPermission, "")
		c.Abort()
	}
}
