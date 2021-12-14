package auth

type relRolePermission struct {
	ID           uint   `json:"id";gorm:"primaryKey"`
	RoleID       string `json:"role_id"`
	PermissionID string `json:"permission_id"`
}

func (relRolePermission) TableName() string {
	return "auth_rel_role_permission"
}
