package database

import (
	"log"
	"os"

	"github.com/jsierrab3991/example_pedidos/libs"
	"github.com/jsierrab3991/example_pedidos/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(enviroment *libs.Enviroment) *gorm.DB {

	db, err := gorm.Open(postgres.Open(enviroment.DatabaseURL), &gorm.Config{})

	if err != nil {
		log.Println("failed to connect database with error: ", err)
		os.Exit(1)
	}
	return db
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&models.User{},
	)
}
