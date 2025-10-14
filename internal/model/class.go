package model

import "gorm.io/gorm"

type Class struct {
	gorm.Model
	Topic          string `gorm:"not null"`
	Description    string `gorm:"not null"`
	GoogleCourseID string `gorm:"null;uniqueIndex"`
	FavScore       int64  `gorm:"default:0"`
	Owner          uint   `gorm:"not null"`
	Status         int    `gorm:"default:1"`
}

func (Class) TableName() string {
	return "class"
}
