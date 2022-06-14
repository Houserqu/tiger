package model

import (
	time "time"

	"gorm.io/plugin/soft_delete"
)

type Role struct {
	ID        uint                  `json:"id" gorm:"primaryKey"`
	Name      string                `json:"name"`
	DeletedAt soft_delete.DeletedAt `json:"deleted_at" swaggertype:"string"`
	CreatedAt time.Time             `json:"created_at"`
	UpdatedAt time.Time             `json:"updated_at"`
}

func (Role) TableName() string {
	return "auth_role"
}
