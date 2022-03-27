package user

import (
	time "time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id";gorm:"primaryKey"`
	Phone     string         `json:"phone"`
	Password  string         `json:"password"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

// TableName 会将 User 的表名重写为 `profiles`
func (User) TableName() string {
	return "user"
}
