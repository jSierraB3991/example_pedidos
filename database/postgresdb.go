package database

import (
	"os"

	"github.com/jsierrab3991/example_pedidos/libs"
	"github.com/jsierrab3991/example_pedidos/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(enviroment *libs.Enviroment) *gorm.DB {

	db, err := gorm.Open(postgres.Open(enviroment.DatabaseURL), &gorm.Config{})

	if err != nil {
		os.Exit(1)
	}
	return db
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&models.User{},
		&models.District{},
		&models.Article{},
		&models.Client{},
		&models.Facture{},
		&models.DetailFacture{},
	)
}
