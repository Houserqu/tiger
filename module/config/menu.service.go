package config

import (
	"errors"

	"github.com/gin-gonic/gin"
	"houserqu.com/gin-starter/core"
)

// 根据权限获取用户的菜单配置
func GetMenusByPermission(c *gin.Context, menus *[]Menu, userId uint) error {
	err := core.Mysql.Find(&menus).Error
	if err != nil {
		core.Log(c).Error(err.Error())
		return errors.New("系统异常")
	}

	// TODO 根据权限过滤菜单

	return nil
}
