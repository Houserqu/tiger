package config

import (
	"github.com/gin-gonic/gin"
	"houserqu.com/gin-starter/core"
)

type PageListDto struct {
	Page    int    `form:"page,default=1"`
	PerPage int    `form:"perPage,default=20"`
	Name    string `form:"name"`
	Path    string `form:"path"`
}

func Controller(r *gin.Engine) {
	// 创建 group 并绑定中间件
	api := r.Group("/api/config")

	api.GET("page", getPage)
	api.GET("pages", getPages)
	api.POST("create-page", createPage)
	api.POST("delete-page", deletePage)
	api.POST("update-page", updatePage)
	api.GET("menus", getMenus)
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

// @Summary 获取页面配置
// @Tags 配置
// @Param path query config.PageListDto true "query"
// @Router /api/config/pages [get]
func getPages(c *gin.Context) {
	var pageListDto PageListDto
	if err := c.ShouldBindQuery(&pageListDto); err != nil {
		core.ResError(c, core.ErrParam, err.Error())
		return
	}

	var pages []Page
	var total int64
	total, err := GetPageList(c, &pages, &pageListDto)
	if err != nil {
		core.ResError(c, core.ErrGetPage, err.Error())
		return
	}

	core.ResSuccess(c, gin.H{
		"items": pages,
		"total": total,
	})
}

// @Summary 获取菜单
// @Tags 配置
// @Router /api/config/menus [get]
func getMenus(c *gin.Context) {
	var menus []Menu
	err := GetMenus(c, &menus)
	if err != nil {
		core.ResError(c, core.ErrGetMenus, err.Error())
		return
	}

	core.ResSuccess(c, menus)
}

type CreatePageDto struct {
	Name   string `json:"name" binding:"required"`
	Path   string `json:"path" binding:"required"`
	Config string `json:"config" binding:"required"`
	Extend string `form:"extend"`
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

type DeletePageDto struct {
	ID uint `form:"id" binding:"required"`
}

// @Summary 删除页面
// @Tags 配置
// @Router /api/config/delete-page [post]
// @Param params body config.DeletePageDto true "参数"
// @Success 200 {number} 1
func deletePage(c *gin.Context) {
	// 参数校验
	var deletePageDto DeletePageDto
	if err := c.ShouldBindJSON(&deletePageDto); err != nil {
		core.ResError(c, core.ErrParam, err.Error())
		return
	}

	id, err := DeletePage(c, deletePageDto.ID)
	if err != nil {
		core.ResError(c, core.ErrDeletePage, err.Error())
		return
	}

	core.ResSuccess(c, id)
}

type UpdatePageDto struct {
	CreatePageDto
	ID uint `form:"id" binding:"required"`
}

// @Summary 更新页面
// @Tags 配置
// @Router /api/config/update-page [post]
// @Param params body config.UpdatePageDto true "参数"
// @Success 200 {number} 1
func updatePage(c *gin.Context) {
	// 参数校验
	var updatePageDto UpdatePageDto
	if err := c.ShouldBindJSON(&updatePageDto); err != nil {
		core.ResError(c, core.ErrParam, err.Error())
		return
	}

	id, err := UpdatePage(c, updatePageDto)
	if err != nil {
		core.ResError(c, core.ErrCreatePage, err.Error())
		return
	}

	core.ResSuccess(c, id)
}
