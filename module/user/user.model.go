package user

import (
	time "time"

	"gorm.io/gorm"
)

type User struct {
	ID       uint           `json:"id";gorm:"primaryKey"`
	Phone    string         `json:"phone"`
	Password string         `json:"password"`
	DeleteAt gorm.DeletedAt `json:"delete_at"`
	CreateAt time.Time      `json:"create_at"`
	UpdateAt time.Time      `json:"update_at"`
}

type Tabler interface {
	TableName() string
}

// TableName 会将 User 的表名重写为 `profiles`
func (User) TableName() string {
	return "auth_user"
}
