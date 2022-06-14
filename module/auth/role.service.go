package auth

import (
	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
	"houserqu.com/tiger/core"
	"houserqu.com/tiger/model"
	"houserqu.com/tiger/utils"
)

//根据id获取角色
func GetRoleById(c *gin.Context, id uint) (role model.Role, err error) {
	err = utils.CRUDReadByID(c, &role, id)
	return role, err
}

func GetRoles(c *gin.Context, roles *[]model.Role) error {
	err := core.Mysql.Find(&roles).Error
	if err != nil {
		core.Log(c).Error(err.Error())
		return errors.New("get role list error")
	}
	return nil
}

//给角色添加权限
func AddPerms(c *gin.Context, roleId uint, permissionIds []string) (relRolePermissions []model.RelRolePermission, err error) {
	//生成所有RolePermission实体
	for _, v := range permissionIds {
		relRolePermission := model.RelRolePermission{
			RoleID:       roleId,
			PermissionID: v,
		}

		relRolePermissions = append(relRolePermissions, relRolePermission)
	}

	// 批量创建（在遇见冲突时，不做任何操作）
	err = core.Mysql.Clauses(clause.OnConflict{DoNothing: true}).Create(&relRolePermissions).Error
	if err != nil {
		return nil, err
	}

	return relRolePermissions, nil

}

// 给角色删除权限
func DelPerms(c *gin.Context, roleId uint, permissionIds []string) (err error) {
	err = core.Mysql.Where("role_id = ? AND permission_id IN ?", roleId, permissionIds).Delete(&model.RelRolePermission{}).Error
	return
}

// 获取角色的所有权限
func GetRolePerms(c *gin.Context, roleId uint) (permissions []model.Permission, err error) {
	err = core.Mysql.Table("auth_rel_role_permission").
		Select("auth_permission.*").
		Joins("left join auth_permission on auth_rel_role_permission.permission_id = auth_permission.id").
		Where("auth_rel_role_permission.role_id = ?", roleId).
		Scan(&permissions).
		Error

	return
}
