package config

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"houserqu.com/gin-starter/core"
)

// @Summary 获取页面配置
// @Tags 配置
// @Param id query int true "page id"
// @Router /api/config/page [get]
func GetPage(c *gin.Context) {
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
