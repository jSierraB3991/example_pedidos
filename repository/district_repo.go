package repository

import (
	"errors"

	"github.com/jsierrab3991/example_pedidos/models"
	"gorm.io/gorm"
)

func (repository Repository) GetDistrict() ([]models.District, error) {
	var districts = []models.District{}
	err := repository.db.Model(&models.District{}).Preload("User").Find(&districts).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("Table of district empty")
		}
		return nil, err
	}
	return districts, nil
}

func (repository Repository) SaveDistrict(district *models.District) (*models.District, error) {
	result := repository.db.Create(&district)
	if result.RowsAffected >= 1 && result.Error != nil {
		return district, nil
	}
	return nil, result.Error
}
