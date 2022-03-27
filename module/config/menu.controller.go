package config

import (
	"github.com/gin-gonic/gin"
	"houserqu.com/gin-starter/core"
)

// @Summary 获取菜单
// @Tags 登录
// @Router /api/config/menus [get]
func GetMenus(c *gin.Context) {
	var menus []Menu
	err := GetMenusByPermission(c, &menus, c.GetUint("userId"))
	if err != nil {
		core.ResError(c, core.ErrGetMenus, err.Error())
		return
	}

	core.ResSuccess(c, menus)
}
