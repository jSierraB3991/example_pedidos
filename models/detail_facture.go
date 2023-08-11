package models

import "gorm.io/gorm"

type DetailFacture struct {
	gorm.Model
	FactureID    uint
	Facture      Facture `gorm:"not null"`
	ArticleID    int
	Article      Article `gorm:"not null"`
	PriceArticle int     `gorm:"column:price_article:not null"`
	Stock        int     `gorm:"not null"`
	Total        int     `gorm:"not null"`
}
