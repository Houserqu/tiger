package menu

import (
	"github.com/gin-gonic/gin"
	"houserqu.com/gin-starter/constants"
	"houserqu.com/gin-starter/core"
	"houserqu.com/gin-starter/middleware"
	"houserqu.com/gin-starter/utils"
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
	api.POST("create", createMenu, middleware.CheckPerm(constants.PER_ADMIN))
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

type CreateMenuReq struct {
	ParentID    uint   `json:"parent_id"`
	Label       string `json:"label" binding:"required"`
	To          string `json:"to"`
	Icon        string `json:"icon"`
	Permissions string `json:"permissions"`
	Target      string `json:"target"`
}

// @Summary 创建菜单
// @Tags 菜单
// @Router /api/menu/create [post]
// @Param params body menu.CreateMenuReq true "参数"
// @Success 200 {number} 1
func createMenu(c *gin.Context) {
	//参数校验
	var createMenuReq CreateMenuReq
	if err := c.ShouldBindJSON(&createMenuReq); err != nil {
		core.ResError(c, constants.ErrParam, err.Error())
		return
	}

	//TODO : check if ParentID exist?

	/*

	 */

	menu := Menu{
		ParentID:    createMenuReq.ParentID,
		Label:       createMenuReq.Label,
		To:          createMenuReq.To,
		Icon:        createMenuReq.Icon,
		Permissions: createMenuReq.Permissions,
		Target:      createMenuReq.Target,
	}

	err := utils.CRUDCreate(c, &menu)
	if err != nil {
		core.ResError(c, constants.ErrCreatePage, err.Error())
		return
	}
	core.ResSuccess(c, menu)
}
