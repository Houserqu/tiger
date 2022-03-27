package config

import time "time"

type Page struct {
	ID        uint      `json:"id";gorm:"primaryKey"`
	Name      string    `json:"name"`
	Path      string    `json:"path"`
	Config    string    `json:"config"`
	Extend    string    `json:"extend"`
	DeletedAt time.Time `json:"deleted_at"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Page) TableName() string {
	return "config_page"
}
