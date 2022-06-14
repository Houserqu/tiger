package model

import (
	time "time"

	"gorm.io/plugin/soft_delete"
)

type User struct {
	ID        uint                  `json:"id" gorm:"primaryKey"`
	Phone     string                `json:"phone"`
	Password  string                `json:"password"`
	DeletedAt soft_delete.DeletedAt `json:"deleted_at" swaggertype:"string"`
	CreatedAt time.Time             `json:"created_at"`
	UpdatedAt time.Time             `json:"updated_at"`
}

// TableName 会将 User 的表名重写为 `profiles`
func (User) TableName() string {
	return "user"
}
