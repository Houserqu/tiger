package role

import (
	"github.com/gin-gonic/gin"
	"houserqu.com/tiger/constants"
	"houserqu.com/tiger/core"
	"houserqu.com/tiger/middleware"
	"houserqu.com/tiger/utils"
)

func Controller(r *gin.Engine) {
	//创建 group 并绑定中间件
	api := r.Group("/api/role", middleware.CheckLogin())

	api.GET("list", getRoleList, middleware.CheckPerm(constants.PER_ADMIN))
	api.POST("create", createRole, middleware.CheckPerm(constants.PER_ADMIN))
	api.POST("delete", deleteMenu, middleware.CheckPerm(constants.PER_ADMIN))
	api.POST("update", updateRole, middleware.CheckPerm(constants.PER_ADMIN))
}

// @Summary 角色列表
// @Tags 角色
// @Router /api/role/list [get]
// @Success 200 {object} role.Role
func getRoleList(c *gin.Context) {
	var roles []Role
	err := GetRoles(c, &roles)

	if err != nil {
		core.ResError(c, constants.ErrGetRoles, err.Error())
		return
	}

	core.ResSuccess(c, roles)
}

type CreateRoleReq struct {
	Name string `json:"name" binding:"required"`
}

// @Summary 创建角色
// @Tags 角色
// @Router /api/role/create [post]
// @Param params body role.CreateRoleReq true "参数"
// @Success 200 {number} 1
func createRole(c *gin.Context) {
	//参数校验
	var createRoleReq CreateRoleReq
	if err := c.ShouldBindJSON(&createRoleReq); err != nil {
		core.ResError(c, constants.ErrParam, err.Error())
		return
	}

	role := Role{
		Name: createRoleReq.Name,
	}

	err := utils.CRUDCreate(c, &role)
	if err != nil {
		core.ResError(c, constants.ErrCreateRole, err.Error())
		return
	}

	core.ResSuccess(c, role)
}

type DeleteRoleReq struct {
	ID uint `form:"id" binding:"required"`
}

// @Summary 删除角色
// @Tags 角色
// @Router /api/role/delete [post]
// @Param params body role.DeleteRoleReq true "参数"
// @Success 200 {number} 1
func deleteMenu(c *gin.Context) {
	//参数校验
	var deleteRoleReq DeleteRoleReq
	if err := c.ShouldBindJSON(&deleteRoleReq); err != nil {
		core.ResError(c, constants.ErrParam, err.Error())
		return
	}

	id, err := utils.CURDDeleteByiD(c, &Role{}, deleteRoleReq.ID)
	if err != nil {
		core.ResError(c, constants.ErrDeleteRole, err.Error())
		return
	}

	core.ResSuccess(c, id)
}

type UpdateRoleReq struct {
	Id   uint   `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

// @Summary 更新角色
// @Tags 角色
// @Router /api/role/update [post]
// @Param params body role.UpdateRoleReq true "参数"
// @Success 200 {number} 1
func updateRole(c *gin.Context) {
	var updateRoleReq map[string]any
	if err := c.ShouldBindJSON(&updateRoleReq); err != nil {
		core.ResError(c, constants.ErrParam, err.Error())
		return
	}

	id, err := utils.CRUDUpdateByID(c, &Role{}, updateRoleReq)
	if err != nil {
		core.ResError(c, constants.ErrUpdateRole, err.Error())
		return
	}

	core.ResSuccess(c, id)

}
