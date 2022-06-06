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
	// api.POST("create", createRole)
	// api.POST("delete", deleteMenu)
	// api.POST("update", updateRole)
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
