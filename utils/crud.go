package utils

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"houserqu.com/gin-starter/core"
)

// 根据 model query 实例查询分页
func CRUDQueryPageList[M any](c *gin.Context, query *gorm.DB, list *[]M) (int64, error) {
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
