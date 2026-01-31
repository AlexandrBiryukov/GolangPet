package userService

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey"`
	Email     string         `gorm:"uniqueIndex;not null"`
	Password  string         `gorm:"not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (User) TableName() string { return "users" }
