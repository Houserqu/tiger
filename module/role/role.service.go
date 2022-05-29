package role

import (
	"errors"

	"github.com/gin-gonic/gin"
	"houserqu.com/tiger/core"
)

func GetRoleByID(id uint) (role Role, err error) {
	err = core.Mysql.First(&role, id).Error
	return
}

func GetRoles(c *gin.Context, roles *[]Role) error {
	err := core.Mysql.Find(&roles).Error
	if err != nil {
		core.Log(c).Error(err.Error())
		return errors.New("c")
	}
	return nil
}
