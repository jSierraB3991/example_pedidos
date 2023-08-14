package models

import (
	"time"

	"gorm.io/gorm"
)

type StateFacture string

const (
	UPDATING StateFacture = "UPDATING"
	FINSIH   StateFacture = "FINISH"
)

type Facture struct {
	gorm.Model
	ID       uint
	Date     time.Time
	ClientID int
	Client   Client `gorm:"not null"`
	UserID   int
	User     User
	Total    int `gorm:"not null"`
	State    StateFacture
}
