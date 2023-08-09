package models

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Name         string `gorm:"not null"`
	Price        int    `gorm:"not null"`
	ModelArticle string `gorm:"not null;column:model"`
	Description  string
}
