package auth

import (
	"github.com/gin-gonic/gin"
	"houserqu.com/tiger/constants"
	"houserqu.com/tiger/constants/PERMISSION"
	"houserqu.com/tiger/core"
	"houserqu.com/tiger/middleware"
	"houserqu.com/tiger/model"
)

func ControllerPermisson(r *gin.Engine) {
	//创建 group 并绑定中间件
	api := r.Group("/api/perm", middleware.CheckLogin(), middleware.CheckPerm(PERMISSION.AUTH_ALL))

	api.GET("list", getPermissionList)
	api.POST("create", createPerm)
	// api.POST("delete", deleteMenu)
	api.POST("update", updatePerm)
	// api.POST("getRoleById", getRoleById)
}

// @Summary 权限列表
// @Tags 权限
// @Router /api/perm/list [get]
// @Success 200 {object} model.Permission
func getPermissionList(c *gin.Context) {
	var perms []model.Permission
	err := GetPerms(c, &perms)

	if err != nil {
		core.ResError(c, constants.ErrGetPerms, err.Error())
		return
	}

	core.ResSuccess(c, perms)
}

type CreatePermReq struct {
	Id   string `json:"id" binding:"required"`
	Desc string `json:"desc" binding:"required"`
	Type string `json:"type" binding:"required"`
}

// @Summary 添加权限
// @Tags 权限
// @Router /api/perm/create [post]
// @Param params body auth.CreatePermReq true "参数"
// @Success 200 {number} 1
func createPerm(c *gin.Context) {
	//参数校验
	var createPermReq CreatePermReq
	if err := c.ShouldBindJSON(&createPermReq); err != nil {
		core.ResError(c, constants.ErrParam, err.Error())
		return
	}

	perm := model.Permission{
		ID:   createPermReq.Id,
		Desc: createPermReq.Desc,
		Type: createPermReq.Type,
	}

	err := CreatePerm(c, &perm)
	if err != nil {
		core.ResError(c, constants.ErrCreatePerm, err.Error())
		return
	}

	core.ResSuccess(c, perm)
}

type UpdatePermReq struct {
	Id   string `json:"id" binding:"required"`
	Desc string `json:"desc"`
	Type string `json:"type"`
}

// @Summary 更新权限
// @Tags 权限
// @Router /api/perm/update [post]
// @Param params body auth.UpdatePermReq true "参数"
// @Success 200 {number} 1
func updatePerm(c *gin.Context) {
	var updatePermReq map[string]any
	if err := c.ShouldBindJSON(&updatePermReq); err != nil {
		core.ResError(c, constants.ErrParam, err.Error())
		return
	}

	id, err := UpdatePermById(c, updatePermReq)
	if err != nil {
		core.ResError(c, constants.ErrUpdatePerm, err.Error())
		return
	}

	core.ResSuccess(c, id)
}
