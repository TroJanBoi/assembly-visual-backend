package model

import "gorm.io/gorm"

type Class struct {
	gorm.Model
	Topic            string `gorm:"not null"`
	Description      string `gorm:"not null"`
	GoogleCourseID   string `gorm:"null;"`
	GoogleCourseLink string `gorm:"null;"`
	GoogleSyncedAt   string `gorm:"null;"`
	FavScore         int64  `gorm:"default:0"`
	Owner            int    `gorm:"not null"`
	Status           int    `gorm:"not null"`
}

func (Class) TableName() string {
	return "class"
}
