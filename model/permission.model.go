package model

import (
	time "time"
)

type Permission struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Desc      string    `json:"desc"`
	Type      string    `json:"type"`
	DeletedAt time.Time `json:"deleted_at"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Permission) TableName() string {
	return "auth_permission"
}
