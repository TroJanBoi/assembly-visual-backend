package model

type BookMark struct {
	ID      uint `gorm:"primaryKey"`
	UserID  uint `gorm:"not null"`
	ClassID uint `gorm:"not null"`
}

func (BookMark) TableName() string {
	return "bookmark"
}
