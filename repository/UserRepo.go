package repository

import (
	"errors"

	"github.com/jsierrab3991/example_pedidos/models"
	"gorm.io/gorm"
)

func (repository *Repository) GetUsers() ([]models.User, error) {
	var users = []models.User{}
	err := repository.db.Find(&users).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("Table of users empty")
		}
		return nil, err
	}
	return users, nil
}
