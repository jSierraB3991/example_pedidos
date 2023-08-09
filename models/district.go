package models

import "gorm.io/gorm"

type District struct {
	gorm.Model
	Name     string `gorm:"not null"`
	Initials string `gorm:"not null"`
	UserID   int
	User     User
}
