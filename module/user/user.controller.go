package user

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"houserqu.com/tiger/constants"
	"houserqu.com/tiger/constants/PERMISSION"
	"houserqu.com/tiger/core"
	"houserqu.com/tiger/middleware"
	"houserqu.com/tiger/model"
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
	api := r.Group("/api/user", middleware.CheckLogin(), middleware.CheckPerm(PERMISSION.USER_ALL))

	api.GET("/:id", GetUser)                                                          // 查单个
	api.GET("/list", middleware.CheckPerm(PERMISSION.USER_LIST), GetUserList)         // 查列表
	api.POST("/create", middleware.CheckPerm(PERMISSION.USER_CREATE), CreateUser)     // 创建
	api.POST("/update", middleware.CheckPerm(PERMISSION.USER_UPDATE), UpdateUser)     // 更新
	api.POST("/delete/:id", middleware.CheckPerm(PERMISSION.USER_DELETE), DeleteUser) // 删除

	api.POST("/addUserRoles", middleware.CheckPerm(PERMISSION.USER_ADD_ROLES), addUserRoles) //添加用户角色
	api.POST("/delUserRoles", middleware.CheckPerm(PERMISSION.USER_DEL_ROLES), delUserRoles) //删除用户角色
}

func GetUser(c *gin.Context) {
	// 参数转换和校验
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		core.ResError(c, constants.ErrParam, "")
		return
	}

	// 根据 ID 查找
	data, err := GetUserByID(uint(id))
	if err != nil {
		core.ResError(c, constants.ErrNotFound, "")
		return
	}

	core.ResSuccess(c, data)
}

func GetUserList(c *gin.Context) {
	var where model.User
	if err := c.ShouldBindQuery(&where); err != nil {
		core.ResError(c, constants.ErrParam, err.Error())
		return
	}

	// 分页参数
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		core.ResError(c, constants.ErrParam, err.Error())
	}
	size, err := strconv.Atoi(c.DefaultQuery("size", "20"))
	if err != nil {
		core.ResError(c, constants.ErrParam, err.Error())
	}

	data, err := GetModelAll(page, size, where)

	if err != nil {
		core.ResError(c, constants.ErrDB, err.Error())
		return
	}

	core.ResSuccess(c, data)
}

// 创建
func CreateUser(c *gin.Context) {
	var params ReqModelCreate
	if err := c.ShouldBindJSON(&params); err != nil {
		core.ResError(c, constants.ErrParam, "")
		return
	}

	result := &model.User{Phone: "123456", Password: "123456"}
	err := CreateModel(result)
	if err != nil {
		core.ResError(c, constants.ErrCreateFail, "")
		return
	}

	core.ResSuccess(c, result)
}

func UpdateUser(c *gin.Context) {
	var params ReqModelUpdate
	if err := c.ShouldBindJSON(&params); err != nil {
		core.ResError(c, constants.ErrParam, "")
		return
	}

	result, err := UpdateModel(params.ID, params.Name, params.Email)
	if err != nil {
		core.ResError(c, constants.ErrUpdateFail, "")
		return
	}

	core.ResSuccess(c, result)
}

func DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		core.ResError(c, constants.ErrParam, "id invalid")
		return
	}

	err = DelModel(id)
	if err != nil {
		core.ResError(c, constants.ErrDeleteFail, "id invalid")
		return
	}

	core.ResSuccess(c, id)
}

type AddUserRolesReq struct {
	UserID  uint   `json:"user_id" binding:"required"`
	RoleIDs []uint `json:"role_id" binding:"required"`
}

// @Summary 添加用户角色
// @Tags 用户
// @Router /api/user/addUserRoles [post]
// @Param params body user.AddUserRolesReq true "参数"
// @Success 200 {object} model.RelUserRole
func addUserRoles(c *gin.Context) {
	var addUserRolesReq AddUserRolesReq
	if err := c.ShouldBindJSON(&addUserRolesReq); err != nil {
		core.ResError(c, constants.ErrParam, err.Error())
		return
	}

	relUserRoles, err := AddUserRoles(c, &addUserRolesReq)
	if err != nil {
		core.ResError(c, constants.ErrAddUserRoles, err.Error())
		return
	}

	core.ResSuccess(c, relUserRoles)
}

type DelUserRolesReq struct {
	UserID  uint   `json:"user_id" binding:"required"`
	RoleIDs []uint `json:"role_id" binding:"required"`
}

// @Summary 为用户移除角色
// @Tags 用户
// @Router /api/user/delUserRoles [post]
// @Param params body user.DelUserRolesReq true "参数"
// @Success 200 {number} 1
func delUserRoles(c *gin.Context) {
	var delUserRolesReq DelUserRolesReq
	if err := c.ShouldBindJSON(&delUserRolesReq); err != nil {
		core.ResError(c, constants.ErrParam, err.Error())
		return
	}

	err := DelUserRoles(c, &delUserRolesReq)
	if err != nil {
		core.ResError(c, constants.ErrDelUserRoles, err.Error())
		return
	}

	core.ResSuccess(c, delUserRolesReq.RoleIDs)
}
