package auth

type RelUserRole struct {
	ID     uint   `json:"id";gorm:"primaryKey"`
	UserID int    `json:"user_id"`
	RoleID string `json:"role_id"`
}

func (RelUserRole) TableName() string {
	return "auth_rel_user_role"
}
