package menu

import (
	"github.com/gin-gonic/gin"
	"houserqu.com/tiger/constants"
	"houserqu.com/tiger/core"
	"houserqu.com/tiger/middleware"
	"houserqu.com/tiger/utils"
)

func Controller(r *gin.Engine) {
	// 创建 group 并绑定中间件
	api := r.Group("/api/menu", middleware.CheckLogin())

	api.GET("list", getMenus, middleware.CheckPerm(constants.PER_ADMIN))
	api.POST("create", createMenu, middleware.CheckPerm(constants.PER_ADMIN))
	api.POST("delete", deleteMenu, middleware.CheckPerm(constants.PER_ADMIN))
	api.POST("update", updateMenu, middleware.CheckPerm(constants.PER_ADMIN))
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
		core.ResError(c, constants.ErrCreateMenu, err.Error())
		return
	}
	core.ResSuccess(c, menu)
}

type DeleteMenuReq struct {
	ID uint `form:"id" binding:"required"`
}

// @Summary 删除菜单
// @Tags 菜单
// @Router /api/menu/delete [post]
// @Param params body menu.DeleteMenuReq true "参数"
// @Success 200 {number} 1
func deleteMenu(c *gin.Context) {
	//参数校验
	var deleteMenuReq DeleteMenuReq
	if err := c.ShouldBindJSON(&deleteMenuReq); err != nil {
		core.ResError(c, constants.ErrParam, err.Error())
		return
	}

	id, err := utils.CURDDeleteByiD(c, &Menu{}, deleteMenuReq.ID)
	if err != nil {
		core.ResError(c, constants.ErrDeleteMenu, err.Error())
		return
	}

	core.ResSuccess(c, id)
}

type UpdateMenuReq struct {
	ID          uint   `json:"id" binding:"required"`
	ParentID    uint   `json:"parent_id"`
	Label       string `json:"label" binding:"required"`
	To          string `json:"to"`
	Icon        string `json:"icon"`
	Permissions string `json:"permissions"`
	Target      string `json:"target"`
}

// @Summary 更新菜单
// @Tags 菜单
// @Router /api/menu/update [post]
// @Param params body menu.UpdateMenuReq true "参数"
// @Success 200 {number} 1
func updateMenu(c *gin.Context) {
	//参数校验
	var updateMenuReq map[string]any
	if err := c.ShouldBindJSON(&updateMenuReq); err != nil {
		core.ResError(c, constants.ErrParam, err.Error())
		return
	}

	id, err := utils.CRUDUpdateByID(c, &Menu{}, updateMenuReq)
	if err != nil {
		core.ResError(c, constants.ErrUpdateMenu, err.Error())
		return
	}

	core.ResSuccess(c, id)
}
