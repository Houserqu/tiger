package auth

import (
	time "time"

	"gorm.io/gorm"
)

type Permission struct {
	ID       string         `json:"id";gorm:"primaryKey"`
	Desc     string         `json:"desc"`
	Type     string         `json:"type"`
	DeleteAt gorm.DeletedAt `json:"delete_at"`
	CreateAt time.Time      `json:"create_at"`
	UpdateAt time.Time      `json:"update_at"`
}

func (Permission) TableName() string {
	return "auth_permission"
}
