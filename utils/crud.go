package utils

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"houserqu.com/tiger/core"
)

// 根据 model query 实例查询分页
func CRUDPageListByQuery[M any](c *gin.Context, query *gorm.DB, list *[]M) (int64, error) {
	// 总数
	var total int64
	query.Count(&total)

	// 处理分页
	query.Scopes(Paginate(c))

	// 开始查询
	err := query.Find(list).Error
	if err != nil {
		core.Log(c).Error(err)
		return total, errors.New("系统异常")
	}

	return total, nil
}

// 根据 model 查询分页列表
func CRUDPageList[M any](c *gin.Context, model *M, list *[]M, where map[string]any) (int64, error) {
	query := core.Mysql.Model(model)

	// 处理查询条件
	for key, value := range where {
		fmt.Println(key, reflect.TypeOf(value).Kind())
		if reflect.TypeOf(value).Kind() == reflect.String && value != "" {
			query = query.Where(key+" LIKE ?", fmt.Sprintf("%%%s%%", value))
		}

		if reflect.TypeOf(value).Kind() == reflect.Int && value != 0 {
			query = query.Where(key+" = ?", value)
		}
	}

	// 总数
	var total int64
	query.Count(&total)

	// 处理分页
	query.Scopes(Paginate(c))

	// 开始查询
	err := query.Find(list).Error
	if err != nil {
		core.Log(c).Error(err)
		return total, errors.New("系统异常")
	}

	return total, nil
}

// 根据 id 查找第一条记录
func CRUDReadByID[M any](c *gin.Context, model *M, id uint) error {
	err := core.Mysql.First(model, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("不存在")
		}

		core.Log(c).Error(err)
		return errors.New("系统异常")
	}

	return nil
}

// 根据 id 删除记录
func CURDDeleteByID[M any](c *gin.Context, model *M, id uint) (uint, error) {
	err := core.Mysql.Delete(model, id).Error
	if err != nil {
		core.Log(c).Error(err)
		return id, errors.New("删除失败")
	}

	return id, nil
}

// 创建记录
func CRUDCreate[M any](c *gin.Context, model *M) error {
	err := core.Mysql.Create(model).Error
	if err != nil {
		core.Log(c).Error(err)
		return errors.New("创建失败")
	}

	return nil
}

// 根据 id 更新记录
func CRUDUpdateByID[M any](c *gin.Context, model *M, params map[string]any) (uint, error) {

	id := uint(params["id"].(float64))

	delete(params, "id")

	res := core.Mysql.Model(model).Where("id = ?", id).Updates(params)
	if res.Error != nil {
		core.Log(c).Error(res.Error)
		return id, errors.New("系统错误")
	}

	if res.RowsAffected == 0 {
		return id, errors.New("未更新")
	}

	return id, nil
}

// 根据 string id 更新记录
func CRUDUpdateByStringID[M any](c *gin.Context, model *M, params map[string]any) (string, error) {

	id := string(params["id"].(string))

	delete(params, "id")

	res := core.Mysql.Model(model).Where("id = ?", id).Updates(params)
	if res.Error != nil {
		core.Log(c).Error(res.Error)
		return id, errors.New("系统错误")
	}

	if res.RowsAffected == 0 {
		return id, errors.New("未更新")
	}

	return id, nil
}
