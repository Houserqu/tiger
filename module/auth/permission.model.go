package auth

import time "time"

type AuthPermission struct {
	ID       string    `json:"id";gorm:"primaryKey"`
	Desc     string    `json:"desc"`
	Type     string    `json:"type"`
	DeleteAt time.Time `json:"delete_at"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}
