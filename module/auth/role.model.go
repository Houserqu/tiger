package auth

import time "time"

type Role struct {
	ID       string    `json:"id";gorm:"primaryKey"`
	Name     string    `json:"name"`
	DeleteAt time.Time `json:"delete_at"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}

func (Role) TableName() string {
	return "auth_role"
}