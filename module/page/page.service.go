package page

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

func GetPageByPath(c *gin.Context, page *Page, path string) error {
	err := core.Mysql.Where(map[string]string{"Path": path}).First(&page).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("页面不存在")
		}

		core.Log(c).Error(err)
		return errors.New("系统异常")
	}

	return nil
}

func GetPageList(c *gin.Context, pages *[]Page, params *PageListReq) (int64, error) {
	query := core.Mysql.Model(&Page{})

	// 处理查询条件
	if params.Name != "" {
		query = query.Where("name LIKE ?", "%"+params.Name+"%")
	}

	if params.Path != "" {
		query = query.Where("path LIKE ?", "%"+params.Path+"%")
	}

	// 总数
	var total int64
	query.Count(&total)

	// 处理分页
	query = query.Limit(params.PerPage).Offset((params.Page - 1) * params.PerPage)

	// 开始查询
	err := query.Find(&pages).Error
	if err != nil {
		core.Log(c).Error(err)
		return total, errors.New("系统异常")
	}

	return total, nil
}

func CreatePage(c *gin.Context, params CreatePageReq) (uint, error) {
	page := Page{Name: params.Name, Path: params.Path, Config: params.Config, Extend: params.Extend, Icon: params.Icon}

	err := core.Mysql.Create(&page).Error
	if err != nil {
		core.Log(c).Error(err)
		return 0, errors.New("创建页面失败")
	}

	return page.ID, nil
}

func UpdatePage(c *gin.Context, params map[string]string) error {

	err := core.Mysql.Model(&Page{}).Where("id = ?", params["id"]).Updates(params).Error
	if err != nil {
		core.Log(c).Error(err)
		return errors.New("系统错误")
	}

	return nil
}

func DeletePage(c *gin.Context, id uint) (uint, error) {
	err := core.Mysql.Delete(&Page{}, id).Error
	if err != nil {
		core.Log(c).Error(err)
		return id, errors.New("删除页面失败")
	}

	return id, nil
}
