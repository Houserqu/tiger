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

// TODO: 多余的，直接在 controller 里调 CURD 方法
func CreateRole(c *gin.Context, role *model.Role) (err error) {

	return utils.CRUDCreate(c, role)
}

// TODO: 多余的，直接在 controller 里调 CURD 方法
//根据id删除角色记录，返回被删除id 与 报错信息
func DeleteRoleById(c *gin.Context, deleteRoleReq *DeleteRoleReq) (uint, error) {
	return utils.CURDDeleteByiD(c, &model.Role{}, deleteRoleReq.ID)
}

// TODO: 多余的，直接在 controller 里调 CURD 方法
//根据id更新角色记录，返回被更新的id 与 报错信息
func UpdateRoleById(c *gin.Context, updateRoleReq map[string]any) (uint, error) {
	return utils.CRUDUpdateByID(c, &model.Role{}, updateRoleReq)
}

//给角色添加权限
func AddPerm(c *gin.Context, addPermReq *AddPermReq) (relRolePermissions []model.RelRolePermission, err error) {

	for _, v := range addPermReq.PermissionIDs {
		relRolePermission := model.RelRolePermission{
			RoleID:       addPermReq.RoleID,
			PermissionID: v,
		}

		relRolePermissions = append(relRolePermissions, relRolePermission)
	}

	for _, relRolePermission := range relRolePermissions {
		//在遇见冲突时，不做任何操作
		err = core.Mysql.Clauses(clause.OnConflict{DoNothing: true}).Create(&relRolePermission).Error
		if err != nil {
			return nil, err
		}
	}
	return relRolePermissions, nil
}

func GetRolePerms(c *gin.Context, getRolePermsReq *GetRolePermsReq, permissions *[]model.Permission) (err error) {
	//查角色权限ID
	var relRolePermission []model.RelRolePermission
	err = core.Mysql.Where("role_id = ?", getRolePermsReq.RoleID).Find(&relRolePermission).Error
	if err != nil {
		return
	}

	//获取permissionId
	var permissionIds []string
	for _, v := range relRolePermission {
		permissionIds = append(permissionIds, v.PermissionID)
	}

	//查权限列表
	err = core.Mysql.Where("id IN ?", permissionIds).Find(permissions).Error
	if err != nil {
		return
	}

	return err
}
