package models

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	ID           uint   `json:"id" mapper:"ID"`
	Name         string `gorm:"not null"`
	Price        int    `gorm:"not null"`
	ModelArticle string `gorm:"not null;column:model"`
	Description  string
}
