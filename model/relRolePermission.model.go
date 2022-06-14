package model

type RelRolePermission struct {
	RoleID       uint   `json:"role_id" gorm:"primaryKey"`
	PermissionID string `json:"permission_id" gorm:"primaryKey"`
}

func (RelRolePermission) TableName() string {
	return "auth_rel_role_permission"
}
