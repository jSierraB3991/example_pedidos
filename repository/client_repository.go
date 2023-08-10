package repository

import (
	"errors"

	"github.com/jsierrab3991/example_pedidos/models"
	"gorm.io/gorm"
)

func (repository Repository) GetClients() ([]models.Client, error) {
	clients := []models.Client{}
	err := repository.db.Model(&models.Client{}).Preload("District").Find(&clients).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("Table of clients empty")
		}
		return nil, err
	}
	return clients, nil
}

func (repository Repository) SaveClient(client *models.Client) (*models.Client, error) {
	result := repository.db.Create(&client)
	if result.RowsAffected >= 1 {
		return client, nil
	}
	return nil, result.Error
}
