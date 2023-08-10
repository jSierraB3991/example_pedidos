package models

import "gorm.io/gorm"

type Client struct {
	gorm.Model
	Name       string `gorm:"not null"`
	LastName   string `gorm:"not null"`
	Telphone   string
	Email      string `gorm:"unique_index;not null"`
	Direction  string `gorm:"not null"`
	DistrictID int
	District   District
}
