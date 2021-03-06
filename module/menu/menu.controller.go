package menu

import (
	"github.com/gin-gonic/gin"
	"houserqu.com/tiger/constants/ERR"
	"houserqu.com/tiger/constants/PERMISSION"
	"houserqu.com/tiger/core"
	"houserqu.com/tiger/middleware"
	"houserqu.com/tiger/model"
	"houserqu.com/tiger/utils"
)

func Controller(r *gin.Engine) {
	// 创建 group 并绑定中间件
	api := r.Group("/api/menu", middleware.CheckLogin(), middleware.CheckPerm(PERMISSION.MENU_ALL))

	api.GET("list", getMenus)
	api.POST("create", createMenu)
	api.POST("delete", deleteMenu)
	api.POST("update", updateMenu)
}

// @Summary 菜单列表
// @Tags 菜单
// @Router /api/menu/list [get]
// @Success 200 {object} model.Menu
func getMenus(c *gin.Context) {
	var menus []model.Menu
	err := GetMenus(c, &menus)
	if err != nil {
		core.ResError(c, ERR.GetMenus, err.Error())
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
		core.ResError(c, ERR.Param, err.Error())
		return
	}

	//TODO : check if ParentID exist?

	/*

	 */

	menu := model.Menu{
		ParentID:    createMenuReq.ParentID,
		Label:       createMenuReq.Label,
		To:          createMenuReq.To,
		Icon:        createMenuReq.Icon,
		Permissions: createMenuReq.Permissions,
		Target:      createMenuReq.Target,
	}

	err := utils.CRUDCreate(c, &menu)
	if err != nil {
		core.ResError(c, ERR.CreateMenu, err.Error())
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
		core.ResError(c, ERR.Param, err.Error())
		return
	}

	id, err := utils.CURDDeleteByID(c, &model.Menu{}, deleteMenuReq.ID)
	if err != nil {
		core.ResError(c, ERR.DeleteMenu, err.Error())
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
		core.ResError(c, ERR.Param, err.Error())
		return
	}

	id, err := utils.CRUDUpdateByID(c, &model.Menu{}, updateMenuReq)
	if err != nil {
		core.ResError(c, ERR.UpdateMenu, err.Error())
		return
	}

	core.ResSuccess(c, id)
}
