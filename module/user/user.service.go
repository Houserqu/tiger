package user

import (
	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"houserqu.com/tiger/core"
	"houserqu.com/tiger/model"
	"houserqu.com/tiger/utils"
)

func GetUserByID(id uint) (user model.User, err error) {
	err = core.Mysql.Take(&user, id).Error
	return
}

func GetModelAll(page int, size int, where interface{}) (user []model.User, err error) {
	if size <= 0 {
		size = 20
	}

	if page <= 0 {
		size = 1
	}

	err = core.Mysql.Where(where).Limit(size).Offset((page - 1) * size).Find(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = nil
		user = []model.User{}
	}
	return
}

func DelModel(id int) (err error) {
	err = core.Mysql.Delete(&model.User{}, id).Error
	return
}

func CreateModel(user *model.User) (err error) {
	err = core.Mysql.Create(&user).Error
	return
}

func UpdateModel(id uint, name string, email string) (user model.User, err error) {
	user, err = GetUserByID(id)
	if err != nil {
		return
	}

	err = core.Mysql.Model(&user).Updates(map[string]interface{}{"name": name, "email": email}).Error
	return
}

func AddUserRoles(c *gin.Context, addUserRolesReq *AddUserRolesReq) (relUserRoles []model.RelUserRole, err error) {
	//生成所有UserRole实体
	for _, v := range addUserRolesReq.RoleIDs {
		relUserRole := model.RelUserRole{
			UserID: addUserRolesReq.UserID,
			RoleID: v,
		}

		relUserRoles = append(relUserRoles, relUserRole)
	}

	//循环添加UserRole
	for _, relUserRole := range relUserRoles {
		//在遇见冲突时，不做任何操作
		err = core.Mysql.Clauses(clause.OnConflict{DoNothing: true}).Create(&relUserRole).Error
		if err != nil {
			return nil, err
		}
	}

	return relUserRoles, nil
}

func DelUserRoles(c *gin.Context, delUserRolesReq *DelUserRolesReq) (err error) {
	//获取要删除权限的实体
	var relUserRoles []model.RelUserRole
	for _, v := range delUserRolesReq.RoleIDs {
		var relUserRole model.RelUserRole
		err = core.Mysql.Where("user_id = ? and role_id = ?", delUserRolesReq.UserID, v).Find(&relUserRole).Error
		if err != nil {
			return err
		}

		relUserRoles = append(relUserRoles, relUserRole)
	}

	for _, v := range relUserRoles {
		_, err = utils.CURDDeleteByID(c, &v, v.ID)
		if err != nil {
			return err
		}
	}
	return nil
}
