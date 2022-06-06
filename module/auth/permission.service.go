package auth

import (
	"errors"

	"github.com/gin-gonic/gin"
	"houserqu.com/tiger/core"
	"houserqu.com/tiger/model"
)

func GetPerms(c *gin.Context, perms *[]model.Permission) error {

	err := core.Mysql.Find(&perms).Error
	if err != nil {
		core.Log(c).Error(err.Error())
		return errors.New("c")
	}
	return nil
}
