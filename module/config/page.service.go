package config

import (
	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"houserqu.com/gin-starter/core"
)

func GetPageByID(c *gin.Context, page *Page, pageId uint) error {
	err := core.Mysql.First(&page, pageId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("页面不存在")
		}

		core.Log(c).Error(err)
		return errors.New("系统异常")
	}

	return nil
}
