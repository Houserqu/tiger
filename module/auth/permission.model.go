package auth

import (
	time "time"

	"gorm.io/gorm"
)

type Permission struct {
	ID        string         `json:"id";gorm:"primaryKey"`
	Desc      string         `json:"desc"`
	Type      string         `json:"type"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

func (Permission) TableName() string {
	return "auth_permission"
}
