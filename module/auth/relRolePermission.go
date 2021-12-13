package auth

type AuthRelRolePermission struct {
	ID           uint   `json:"id";gorm:"primaryKey"`
	RoleID       string `json:"role_id"`
	PermissionID string `json:"permission_id"`
}
