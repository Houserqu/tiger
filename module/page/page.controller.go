package page

import (
	"github.com/gin-gonic/gin"
	"houserqu.com/gin-starter/constants"
	"houserqu.com/gin-starter/core"
	"houserqu.com/gin-starter/middleware"
)

type PageListReq struct {
	Page    int    `form:"page,default=1"`
	PerPage int    `form:"perPage,default=20"`
	Name    string `form:"name"`
	Path    string `form:"path"`
}

func Controller(r *gin.Engine) {
	// 创建 group 并绑定中间件
	api := r.Group("/api/page", middleware.CheckLogin())

	api.GET("detail", getPage)
	api.GET("list", getPages)
	api.POST("create", createPage)
	api.POST("delete", deletePage)
	api.POST("update", updatePage)
}

// @Summary 页面详情
// @Tags 页面
// @Param path query string true "path"
// @Router /api/config/page [get]
func getPage(c *gin.Context) {
	var page Page
	err := GetPageByPath(c, &page, c.Query("path"))
	if err != nil {
		core.ResError(c, constants.ErrGetPage, err.Error())
		return
	}

	core.ResSuccess(c, page)
}

// @Summary 页面列表
// @Tags 页面
// @Param path query page.PageListReq true "query"
// @Router /api/page/list [get]
func getPages(c *gin.Context) {
	var pageListDto PageListReq
	if err := c.ShouldBindQuery(&pageListDto); err != nil {
		core.ResError(c, constants.ErrParam, err.Error())
		return
	}

	var pages []Page
	var total int64
	total, err := GetPageList(c, &pages, &pageListDto)
	if err != nil {
		core.ResError(c, constants.ErrGetPage, err.Error())
		return
	}

	core.ResSuccess(c, gin.H{
		"items": pages,
		"total": total,
	})
}

type CreatePageReq struct {
	Name   string `json:"name" binding:"required"`
	Path   string `json:"path" binding:"required"`
	Icon   string `json:"icon"`
	Config string `json:"config" binding:"required"`
	Extend string `form:"extend"`
}

// @Summary 创建页面
// @Tags 页面
// @Router /api/page/create [post]
// @Param params body page.CreatePageReq true "参数"
// @Success 200 {number} 1
func createPage(c *gin.Context) {
	// 参数校验
	var createPageReq CreatePageReq
	if err := c.ShouldBindJSON(&createPageReq); err != nil {
		core.ResError(c, constants.ErrParam, err.Error())
		return
	}

	id, err := CreatePage(c, createPageReq)
	if err != nil {
		core.ResError(c, constants.ErrCreatePage, err.Error())
		return
	}

	core.ResSuccess(c, id)
}

type DeletePageReq struct {
	ID uint `form:"id" binding:"required"`
}

// @Summary 删除页面
// @Tags 页面
// @Router /api/page/delete [post]
// @Param params body page.DeletePageReq true "参数"
// @Success 200 {number} 1
func deletePage(c *gin.Context) {
	// 参数校验
	var deletePageReq DeletePageReq
	if err := c.ShouldBindJSON(&deletePageReq); err != nil {
		core.ResError(c, constants.ErrParam, err.Error())
		return
	}

	id, err := DeletePage(c, deletePageReq.ID)
	if err != nil {
		core.ResError(c, constants.ErrDeletePage, err.Error())
		return
	}

	core.ResSuccess(c, id)
}

// @Summary 更新页面
// @Tags 页面
// @Router /api/page/update [post]
// @Param params body map[string]interface{} true "参数"
// @Success 200 {number} 1
func updatePage(c *gin.Context) {
	// 参数校验
	var updatePageReq map[string]string
	if err := c.ShouldBindJSON(&updatePageReq); err != nil {
		core.ResError(c, constants.ErrParam, err.Error())
		return
	}

	err := UpdatePage(c, updatePageReq)
	if err != nil {
		core.ResError(c, constants.ErrUpdatePage, err.Error())
		return
	}

	core.ResSuccess(c, updatePageReq["id"])
}
