package model

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Notification struct {
	gorm.Model
	UserID  int            `gorm:"not null"`
	Type    string         `gorm:"not null"`
	Title   string         `gorm:"not null"`
	Message string         `gorm:"not null"`
	Data    datatypes.JSON `gorm:"type:jsonb;null"`
	IsRead  bool           `gorm:"not null;default:false"`
	User    User           `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (Notification) TableName() string {
	return "notification"
}
