package auth

type AuthRelUserRole struct {
	ID     uint   `json:"id";gorm:"primaryKey"`
	UserID int    `json:"user_id"`
	RoleID string `json:"role_id"`
}
