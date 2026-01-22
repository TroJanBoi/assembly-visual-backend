package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email        string `gorm:"uniqueIndex;not null"`
	PasswordHash string `gorm:"null"`
	Name         string `gorm:"null"`
	PicturePath  string `gorm:"null"`
}

func (User) TableName() string {
	return "user"
}
