package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string `gorm:"null"`
	Email     string `gorm:"uniqueIndex;not null"`
	Password  string `gorm:"not null"`
	Is_active bool   `gorm:"default:true"`
	Is_admin  bool   `gorm:"default:false"`
	Name      string `gorm:"null"`
	Tel       string `gorm:"null"`
}

func (User) TableName() string {
	return "users"
}
