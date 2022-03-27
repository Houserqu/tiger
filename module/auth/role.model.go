package auth

import time "time"

type Role struct {
	ID        string    `json:"id";gorm:"primaryKey"`
	Name      string    `json:"name"`
	DeletedAt time.Time `json:"deleted_at"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Role) TableName() string {
	return "auth_role"
}
