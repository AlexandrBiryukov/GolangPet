package userService

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	ID     uint `gorm:"primaryKey"`
	Task   string
	IsDone bool
	UserID uint
}

func (Task) TableName() string { return "tasks" }

type User struct {
	ID        uint           `gorm:"primaryKey"`
	Email     string         `gorm:"uniqueIndex;not null"`
	Password  string         `gorm:"not null"`
	Tasks     []Task         `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (User) TableName() string { return "users" }
