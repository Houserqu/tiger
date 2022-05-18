package menu

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"houserqu.com/gin-starter/core"
)

// 获取所有菜单
func GetMenus(c *gin.Context, menus *[]Menu) error {
	err := core.Mysql.Find(&menus).Error
	if err != nil {
		core.Log(c).Error(err.Error())
		return errors.New("c")
	}

	return nil
}

// 根据权限获取用户的菜单配置
func GetMenusByUserId(c *gin.Context, menus *[]Menu, userId uint) error {
	err := core.Mysql.Find(&menus).Error
	if err != nil {
		core.Log(c).Error(fmt.Sprintf("%s; userId = %d", err, userId))
		return errors.New("查找用户菜单失败")
	}

	// TODO 根据权限过滤菜单

	return nil
}
