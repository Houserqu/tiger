package page

import (
	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"houserqu.com/gin-starter/core"
)

func GetPageByPath(c *gin.Context, path string) (page Page, err error) {
	err = core.Mysql.Where(map[string]string{"Path": path}).First(&page).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("页面不存在")
			return
		}

		core.Log(c).Error(err)
		err = errors.New("系统异常")
		return
	}

	return
}
