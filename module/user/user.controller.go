package user

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"houserqu.com/gin-starter/constants"
	"houserqu.com/gin-starter/core"
	"houserqu.com/gin-starter/middleware"
)

type ReqModelCreate struct {
	Name  string `form:"name" binding:"required"`
	Email string `form:"email" binding:"required"`
	Age   int    `form:"age" binding:"required"`
}

type ReqModelUpdate struct {
	ID    uint   `form:"id" binding:"required"`
	Name  string `form:"name" binding:"required"`
	Email string `form:"email" binding:"required"`
}

func Controller(r *gin.Engine) {
	// 创建 group 并绑定中间件
	api := r.Group("/api/user", middleware.CheckLogin(), middleware.CheckPerm(constants.PER_USER_ALL))

	api.GET("/:id", GetUser)                                                             // 查单个
	api.GET("/list", middleware.CheckPerm(constants.PER_USER_LIST), GetUserList)         // 查列表
	api.POST("/create", middleware.CheckPerm(constants.PER_USER_CREATE), CreateUser)     // 创建
	api.POST("/update", middleware.CheckPerm(constants.PER_USER_UPDATE), UpdateUser)     // 更新
	api.POST("/delete/:id", middleware.CheckPerm(constants.PER_USER_DELETE), DeleteUser) // 删除
}

func GetUser(c *gin.Context) {
	// 参数转换和校验
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		core.ResError(c, core.ErrParam, "")
		return
	}

	// 根据 ID 查找
	data, err := GetUserByID(uint(id))
	if err != nil {
		core.ResError(c, core.ErrNotFound, "")
		return
	}

	core.ResSuccess(c, data)
}

func GetUserList(c *gin.Context) {
	var where User
	if err := c.ShouldBindQuery(&where); err != nil {
		core.ResError(c, core.ErrParam, err.Error())
		return
	}

	// 分页参数
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, err := strconv.Atoi(c.DefaultQuery("size", "20"))
	if err != nil {
		core.ResError(c, core.ErrParam, err.Error())
	}

	data, err := GetModelAll(page, size, where)

	if err != nil {
		core.ResError(c, core.ErrDB, err.Error())
		return
	}

	core.ResSuccess(c, data)
}

// 创建
func CreateUser(c *gin.Context) {
	var params ReqModelCreate
	if err := c.ShouldBindJSON(&params); err != nil {
		core.ResError(c, core.ErrParam, "")
		return
	}

	result := &User{Phone: "123456", Password: "123456"}
	err := CreateModel(result)
	if err != nil {
		core.ResError(c, core.ErrCreateFail, "")
		return
	}

	core.ResSuccess(c, result)
}

func UpdateUser(c *gin.Context) {
	var params ReqModelUpdate
	if err := c.ShouldBindJSON(&params); err != nil {
		core.ResError(c, core.ErrParam, "")
		return
	}

	result, err := UpdateModel(params.ID, params.Name, params.Email)
	if err != nil {
		core.ResError(c, core.ErrUpdateFail, "")
		return
	}

	core.ResSuccess(c, result)
}

func DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		core.ResError(c, core.ErrParam, "id invalid")
		return
	}

	err = DelModel(id)
	if err != nil {
		core.ResError(c, core.ErrDeleteFail, "id invalid")
		return
	}

	core.ResSuccess(c, id)
}
