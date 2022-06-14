package model

import (
	time "time"

	"gorm.io/plugin/soft_delete"
)

type Page struct {
	ID        uint                  `json:"id" gorm:"primaryKey"`
	Name      string                `json:"name"`
	Path      string                `json:"path"`
	Icon      string                `json:"icon"`
	Config    string                `json:"config"`
	Extend    string                `json:"extend"`
	DeletedAt soft_delete.DeletedAt `json:"deleted_at" swaggertype:"string"`
	CreatedAt time.Time             `json:"created_at"`
	UpdatedAt time.Time             `json:"updated_at"`
}

func (Page) TableName() string {
	return "config_page"
}
