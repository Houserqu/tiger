package login

import (
	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"houserqu.com/gin-starter/core"
	"houserqu.com/gin-starter/module/user"
)

// 根据手机号+密码校验用户
func CheckUserByPhoneAndPassword(c *gin.Context, phone string, password string) (user.User, error) {
	var user user.User
	err := core.Mysql.Where(map[string]any{"Phone": phone}).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, errors.New("手机号不存在")
		}

		// 不在预期内的异常记录日志，不能直接暴露给
		core.Log(c).Error(err)
		return user, errors.New("系统异常")
	}

	if user.Password != password {
		err = errors.New("密码错误")
	}

	return user, nil
}
