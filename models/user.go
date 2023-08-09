package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"not null"`
	LastName string `gorm:"not null"`
	Telphone string
	Email    string `gorm:"unique_index;not null"`
	Password string `gorm:"not null"`
}
