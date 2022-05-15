package menu

import (
	"github.com/gin-gonic/gin"
	"houserqu.com/gin-starter/constants"
	"houserqu.com/gin-starter/core"
	"houserqu.com/gin-starter/middleware"
)

type PageListDto struct {
	Page    int    `form:"page,default=1"`
	PerPage int    `form:"perPage,default=20"`
	Name    string `form:"name"`
	Path    string `form:"path"`
}

func Controller(r *gin.Engine) {
	// 创建 group 并绑定中间件
	api := r.Group("/api/menu", middleware.CheckLogin())

	api.GET("list", getMenus, middleware.CheckPerm(constants.PER_ADMIN))
}

// @Summary 菜单列表
// @Tags 菜单
// @Router /api/menu/list [get]
// @Success 200 {object} menu.Menu
func getMenus(c *gin.Context) {
	var menus []Menu
	err := GetMenus(c, &menus)
	if err != nil {
		core.ResError(c, constants.ErrGetMenus, err.Error())
		return
	}

	core.ResSuccess(c, menus)
}