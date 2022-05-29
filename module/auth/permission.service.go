package auth

import (
	"houserqu.com/tiger/core"
	"houserqu.com/tiger/model"
)

//已经迁移到middleware.checkPerm中，感觉可以删掉
func GetUserPermissions(userId uint) (allowPermIds []string, err error) {
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
