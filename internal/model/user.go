package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
	Name     string `gorm:"null"`
	Tel      string `gorm:"null"`
}

func (User) TableName() string {
	return "users"
}
