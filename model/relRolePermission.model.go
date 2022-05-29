package model

type RelRolePermission struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	RoleID       string `json:"role_id"`
	PermissionID string `json:"permission_id"`
}

func (RelRolePermission) TableName() string {
	return "auth_rel_role_permission"
}
