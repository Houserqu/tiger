package config

import (
	"github.com/gin-gonic/gin"
	"houserqu.com/gin-starter/core"
	"houserqu.com/gin-starter/middleware"
)

func Controller(r *gin.Engine) {
	// 创建 group 并绑定中间件
	api := r.Group("/api/config", middleware.CheckLogin())

	api.GET("page", getPage)
	api.GET("menus", getMenus)
	api.POST("create-page", createPage)
}

// @Summary 获取页面配置
// @Tags 配置
// @Param path query string true "path"
// @Router /api/config/page [get]
func getPage(c *gin.Context) {
	var page Page
	err := GetPageByPath(c, &page, c.Query("path"))
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

type CreatePageDto struct {
	Name   string                 `json:"name" binding:"required"`
	Path   string                 `json:"path" binding:"required"`
	Config map[string]interface{} `json:"config" binding:"required"`
	Extend map[string]interface{} `json:"extend" binding:"required"`
}

// @Summary 创建页面
// @Tags 配置
// @Router /api/config/create-page [post]
// @Param params body config.CreatePageDto true "参数"
// @Success 200 {number} 1
func createPage(c *gin.Context) {
	// 参数校验
	var createPagetDto CreatePageDto
	if err := c.ShouldBindJSON(&createPagetDto); err != nil {
		core.ResError(c, core.ErrParam, err.Error())
		return
	}

	id, err := CreatePage(c, createPagetDto)
	if err != nil {
		core.ResError(c, core.ErrCreatePage, err.Error())
		return
	}

	core.ResSuccess(c, id)
}
