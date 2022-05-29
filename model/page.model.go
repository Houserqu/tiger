package model

import (
	time "time"

	"gorm.io/gorm"
)

type Page struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name"`
	Path      string         `json:"path"`
	Icon      string         `json:"icon"`
	Config    string         `json:"config"`
	Extend    string         `json:"extend"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

func (Page) TableName() string {
	return "config_page"
}
