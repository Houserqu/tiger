package page

import (
	"github.com/gin-gonic/gin"
	"houserqu.com/tiger/constants/ERR"
	"houserqu.com/tiger/core"
	"houserqu.com/tiger/middleware"
	"houserqu.com/tiger/model"
	"houserqu.com/tiger/utils"
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
	page, err := GetPageByPath(c, c.Query("path"))
	if err != nil {
		core.ResError(c, ERR.GetPage, err.Error())
		return
	}

	core.ResSuccess(c, page)
}

// @Summary 页面列表
// @Tags 页面
// @Param path query page.PageListReq true "query"
// @Router /api/page/list [get]
func getPages(c *gin.Context) {
	var pageListReq PageListReq
	if err := c.ShouldBindQuery(&pageListReq); err != nil {
		core.ResError(c, ERR.Param, err.Error())
		return
	}

	where := map[string]any{
		"name": pageListReq.Name,
		"path": pageListReq.Path,
	}

	var pages []model.Page
	total, err := utils.CRUDPageList(c, &model.Page{}, &pages, where)
	if err != nil {
		core.ResError(c, ERR.GetPage, err.Error())
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
		core.ResError(c, ERR.Param, err.Error())
		return
	}

	page := model.Page{
		Name:   createPageReq.Name,
		Path:   createPageReq.Path,
		Config: createPageReq.Config,
		Extend: createPageReq.Extend,
		Icon:   createPageReq.Icon,
	}

	err := utils.CRUDCreate(c, &page)
	if err != nil {
		core.ResError(c, ERR.CreatePage, err.Error())
		return
	}

	core.ResSuccess(c, page)
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
		core.ResError(c, ERR.Param, err.Error())
		return
	}

	id, err := utils.CURDDeleteByID(c, &model.Page{}, deletePageReq.ID)
	if err != nil {
		core.ResError(c, ERR.DeletePage, err.Error())
		return
	}

	core.ResSuccess(c, id)
}

// @Summary 更新页面
// @Tags 页面
// @Router /api/page/update [post]
// @Param params body page.UpdatePageReq true "参数"
// @Success 200 {number} 1
func updatePage(c *gin.Context) {
	// 参数校验
	var updatePageReq map[string]any
	c.ShouldBindJSON(&updatePageReq)

	id, err := utils.CRUDUpdateByID(c, &model.Page{}, updatePageReq)
	if err != nil {
		core.ResError(c, ERR.UpdatePage, err.Error())
		return
	}

	core.ResSuccess(c, id)
}
