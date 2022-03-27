package config

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"houserqu.com/gin-starter/core"
	"houserqu.com/gin-starter/middleware"
)

func Controller(r *gin.Engine) {
	// 创建 group 并绑定中间件
	api := r.Group("/api/config", middleware.CheckLogin())

	api.GET("page", getPage)
	api.GET("menus", getMenus)
}

// @Summary 获取页面配置
// @Tags 配置
// @Param id query int true "page id"
// @Router /api/config/page [get]
func getPage(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		core.ResError(c, core.ErrParam, err.Error())
		return
	}

	var page Page
	err = GetPageByID(c, &page, uint(id))
	if err != nil {
		core.ResError(c, core.ErrGetPage, err.Error())
		return
	}

	core.ResSuccess(c, page)
}

// @Summary 获取菜单
// @Tags 配置
// @Router /api/config/menus [get]
func getMenus(c *gin.Context) {
	var menus []Menu
	err := GetMenusByPermission(c, &menus, c.GetUint("userId"))
	if err != nil {
		core.ResError(c, core.ErrGetMenus, err.Error())
		return
	}

	core.ResSuccess(c, menus)
}
