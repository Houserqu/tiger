package model

import (
	time "time"

	"gorm.io/plugin/soft_delete"
)

type Permission struct {
	ID        string                `json:"id" gorm:"primaryKey"`
	Desc      string                `json:"desc"`
	Type      string                `json:"type"`
	DeletedAt soft_delete.DeletedAt `json:"deleted_at" swaggertype:"string"`
	CreatedAt time.Time             `json:"created_at"`
	UpdatedAt time.Time             `json:"updated_at"`
}

func (Permission) TableName() string {
	return "auth_permission"
}
