package model

import (
	time "time"

	"gorm.io/gorm"
)

type Menu struct {
	ID          uint           `json:"id"`
	ParentID    uint           `json:"parent_id"`
	Label       string         `json:"label"`
	To          string         `json:"to"`
	Icon        string         `json:"icon"`
	Permissions string         `json:"permissions"`
	Target      string         `json:"target"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" swaggertype:"string"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}

func (Menu) TableName() string {
	return "config_menu"
}
