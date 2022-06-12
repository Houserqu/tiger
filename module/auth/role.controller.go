package auth

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"houserqu.com/tiger/constants"
	"houserqu.com/tiger/constants/PERMISSION"
	"houserqu.com/tiger/core"
	"houserqu.com/tiger/middleware"
	"houserqu.com/tiger/model"
)

func Controller(r *gin.Engine) {
	//创建 group 并绑定中间件
	api := r.Group("/api/role", middleware.CheckLogin(), middleware.CheckPerm(PERMISSION.AUTH_ALL))

	//角色操作相关接口
	api.GET("list", getRoleList)
	api.POST("create", createRole)
	api.POST("delete", deleteRole)
	api.POST("update", updateRole)
	api.GET("detail/:id", getRoleById)

	//为角色权限操作
	api.POST("addPerms", addPerms)
	api.POST("delPerms", delPerms)
	api.POST("getRolePerms", getRolePerms)
}

type GetRoleReq struct {
	ID uint `form:"id" binding:"required"`
}

// @Summary 根据id获取角色
// @Tags 角色
// @Router /api/role/detail/{id} [get]
// @Param id path int true "角色 ID"
// @Success 200 {object} model.Role
func getRoleById(c *gin.Context) {
	//参数校验
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		core.ResError(c, constants.ErrParam, "")
		return
	}

	role, err := GetRoleById(c, uint(id))
	if err != nil {
		core.ResError(c, constants.ErrGetRole, err.Error())
		return
	}

	core.ResSuccess(c, role)
}

// @Summary 角色列表
// @Tags 角色
// @Router /api/role/list [get]
// @Success 200 {object} model.Role
func getRoleList(c *gin.Context) {
	var roles []model.Role
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
// @Param params body auth.CreateRoleReq true "参数"
// @Success 200 {number} 1
func createRole(c *gin.Context) {
	//参数校验
	var createRoleReq CreateRoleReq
	if err := c.ShouldBindJSON(&createRoleReq); err != nil {
		core.ResError(c, constants.ErrParam, err.Error())
		return
	}

	role := model.Role{
		Name: createRoleReq.Name,
	}

	err := CreateRole(c, &role)
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
// @Param params body auth.DeleteRoleReq true "参数"
// @Success 200 {number} 1
func deleteRole(c *gin.Context) {
	//参数校验
	var deleteRoleReq DeleteRoleReq
	if err := c.ShouldBindJSON(&deleteRoleReq); err != nil {
		core.ResError(c, constants.ErrParam, err.Error())
		return
	}

	id, err := DeleteRoleById(c, &deleteRoleReq)
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
// @Param params body auth.UpdateRoleReq true "参数"
// @Success 200 {number} 1
func updateRole(c *gin.Context) {
	var updateRoleReq map[string]any
	if err := c.ShouldBindJSON(&updateRoleReq); err != nil {
		core.ResError(c, constants.ErrParam, err.Error())
		return
	}

	id, err := UpdateRoleById(c, updateRoleReq)
	if err != nil {
		core.ResError(c, constants.ErrUpdateRole, err.Error())
		return
	}

	core.ResSuccess(c, id)
}

type AddPermsReq struct {
	RoleID        string   `json:"role_id" binding:"required"`
	PermissionIDs []string `json:"permission_ids" binding:"required"`
}

// @Summary 为角色添加权限
// @Tags 角色
// @Router /api/role/addPerms [post]
// @Param params body auth.AddPermsReq true "参数"
// @Success 200 {number} 1
func addPerms(c *gin.Context) {
	var addPermsReq AddPermsReq
	if err := c.ShouldBindJSON(&addPermsReq); err != nil {
		core.ResError(c, constants.ErrParam, err.Error())
		return
	}

	relRolePermissions, err := AddPerms(c, &addPermsReq)
	if err != nil {
		core.ResError(c, constants.ErrAddPerm, err.Error())
		return
	}

	core.ResSuccess(c, relRolePermissions)
}

type DelPermsReq struct {
	RoleID        string   `json:"role_id" binding:"required"`
	PermissionIDs []string `json:"permission_ids" binding:"required"`
}

// @Summary 为角色移除权限
// @Tags 角色
// @Router /api/role/delPerms [post]
// @Param params body auth.DelPermsReq true "参数"
// @Success 200 {number} 1
func delPerms(c *gin.Context) {
	var delPermsReq DelPermsReq
	if err := c.ShouldBindJSON(&delPermsReq); err != nil {
		core.ResError(c, constants.ErrParam, err.Error())
		return
	}

	err := DelPerms(c, &delPermsReq)
	if err != nil {
		core.ResError(c, constants.ErrDelPerms, err.Error())
		return
	}

	core.ResSuccess(c, delPermsReq.PermissionIDs)
}

type GetRolePermsReq struct {
	RoleID string `json:"role_id" binding:"required"`
}

// @Summary 根据角色id获取角色权限
// @Tags 角色
// @Router /api/role/getRolePerms [post]
// @Param params body auth.GetRolePermsReq true "参数"
// @Success 200 {object} model.Permission
func getRolePerms(c *gin.Context) {
	var rolePerms []model.Permission
	var getRolePermsReq GetRolePermsReq
	if err := c.ShouldBindJSON(&getRolePermsReq); err != nil {
		core.ResError(c, constants.ErrParam, err.Error())
		return
	}

	err := GetRolePerms(c, &getRolePermsReq, &rolePerms)
	if err != nil {
		core.ResError(c, constants.ErrGetRolePerms, err.Error())
		return
	}

	core.ResSuccess(c, rolePerms)
}