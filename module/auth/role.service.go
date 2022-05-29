package auth

import (
	"errors"

	"github.com/gin-gonic/gin"
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
		return errors.New("c")
	}
	return nil
}

func CreateRole(c *gin.Context, role *model.Role) (err error) {

	return utils.CRUDCreate(c, role)
}

//根据id删除角色记录，返回被删除id 与 报错信息
func DeleteRoleById(c *gin.Context, deleteRoleReq *DeleteRoleReq) (uint, error) {
	return utils.CURDDeleteByiD(c, &model.Role{}, deleteRoleReq.ID)
}

//根据id更新角色记录，返回被更新的id 与 报错信息
func UpdateRoleById(c *gin.Context, updateRoleReq map[string]any) (uint, error) {
	return utils.CRUDUpdateByID(c, &model.Role{}, updateRoleReq)
}
