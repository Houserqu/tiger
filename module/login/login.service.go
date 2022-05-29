package login

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"houserqu.com/tiger/core"
	"houserqu.com/tiger/model"
	"houserqu.com/tiger/module/menu"
)

type UserInfo struct {
	ID    uint   `json:"id"`
	Phone string `json:"phone"`
}
type AdminLoginInfo struct {
	UserInfo UserInfo     `json:"userInfo"`
	Menus    []model.Menu `json:"menus"`
}

// 根据手机号+密码校验用户
func CheckUserByPhoneAndPassword(c *gin.Context, phone string, password string) (model.User, error) {
	var user model.User
	err := core.Mysql.Where(map[string]string{"Phone": phone}).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, errors.New("手机号不存在")
		}

		// 不在预期内的异常记录日志，不能直接暴露给前端
		core.Log(c).Error(err)
		return user, errors.New("DB 错误")
	}

	if user.Password != password {
		return user, errors.New("密码错误")
	}

	return user, nil
}

// 获取管理台用户登录信息
func GetAdminLoginInfoByUserId(c *gin.Context, adminLoginInfo *AdminLoginInfo, userId uint) error {
	// 查用户基本信息
	var user model.User
	err := core.Mysql.First(&user, userId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("手机号不存在")
		}

		// TODO: 输出 sql
		core.Log(c).Error(err)
		return errors.New("查询用户信息失败")
	}

	adminLoginInfo.UserInfo.ID = user.ID
	adminLoginInfo.UserInfo.Phone = user.Phone

	// TODO: 查权限信息

	// 查菜单信息
	err = menu.GetMenusByUserId(c, &adminLoginInfo.Menus, userId)

	return err
}
