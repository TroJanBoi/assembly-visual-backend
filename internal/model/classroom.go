package model

import "gorm.io/gorm"

type Classroom struct {
	gorm.Model
	OwnerId          int    `gorm:"not null"`
	Topic            string `gorm:"not null"`
	Description      string `gorm:"null"`
	Status           int    `gorm:"not null;default:0"` // 0: public, 1: private, 2: archived
	GoogleCourseID   string `gorm:"null;"`
	GoogleCourseLink string `gorm:"null;"`
	GoogleSyncedAt   string `gorm:"null;"`
	User             User   `gorm:"foreignKey:Owner_id;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (Classroom) TableName() string {
	return "classroom"
}
