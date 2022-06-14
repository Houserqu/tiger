package model

type RelUserRole struct {
	UserID uint `json:"user_id" gorm:"primaryKey"`
	RoleID uint `json:"role_id" gorm:"primaryKey"`
}

func (RelUserRole) TableName() string {
	return "auth_rel_user_role"
}
