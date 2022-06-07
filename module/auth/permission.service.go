package auth

import (
	"errors"

	"github.com/gin-gonic/gin"
	"houserqu.com/tiger/core"
	"houserqu.com/tiger/model"
	"houserqu.com/tiger/utils"
)

func GetPerms(c *gin.Context, perms *[]model.Permission) error {

	err := core.Mysql.Find(&perms).Error
	if err != nil {
		core.Log(c).Error(err.Error())
		return errors.New("c")
	}
	return nil
}

func CreatePerm(c *gin.Context, perm *model.Permission) (err error) {
	return utils.CRUDCreate(c, perm)
}

func UpdatePermById(c *gin.Context, updatePermReq map[string]any) (uint, error) {
	return utils.CRUDUpdateByID(c, &model.Permission{}, updatePermReq)
}
