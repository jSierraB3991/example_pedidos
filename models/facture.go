package models

import (
	"time"

	"gorm.io/gorm"
)

type Facture struct {
	gorm.Model
	Date     time.Time
	ClientID int
	Client   Client `gorm:"not null"`
	UserID   int
	User     User
	Total    int `gorm:"not null"`
}
