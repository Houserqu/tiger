package example

import (
	"github.com/gin-gonic/gin"
	"houserqu.com/tiger/constants/ERR"
	"houserqu.com/tiger/core"
)

type ReqModelCreate struct {
	Name  string `form:"name" binding:"required"`
	Email string `form:"email" binding:"required"`
	Age   int    `form:"age" binding:"required"`
}

type ReqModelUpdate struct {
	ID    int    `form:"id" binding:"required"`
	Name  string `form:"name" binding:"required"`
	Email string `form:"email" binding:"required"`
}

func GetModel(c *gin.Context) {
	core.Log(c).Info("example")
	// 根据 ID 查找
	data, err := GetModelByID()
	if err != nil {
		core.ResErrorWithData(c, ERR.NotFound, "model 不存在", err.Error())
		return
	}

	core.ResSuccess(c, data)
}

func ErrorExample(c *gin.Context) {
	// 根据 ID 查找
	core.ResError(c, ERR.CreateFail, "用户失败了")
}
