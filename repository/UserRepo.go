package repository

import (
	"errors"
	"fmt"

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

func (repository *Repository) SaveUser(user *models.User) (*models.User, error) {
	result := repository.db.Create(&user)
	if result.RowsAffected >= 1 && result.Error != nil {
		return user, nil
	}
	return nil, result.Error
}

func (repository *Repository) UpdateUser(user *models.User, id string) (*models.User, error) {
	exists, err := repository.FindById(id)
	if err != nil {
		return user, err
	}
	user.ID = exists.ID
	result := repository.db.Model(&exists).Updates(models.User{Email: user.Email, Password: user.Password})
	if result.RowsAffected >= 1 && result.Error != nil {
		return user, nil
	}
	return nil, result.Error
}

func (repository *Repository) FindById(id string) (*models.User, error) {
	var exists models.User
	repository.db.Find(&exists, id)
	if exists.ID == 0 {
		return nil, errors.New(fmt.Sprintf("User witch id: %v Not Exists", id))
	}
	return &exists, nil
}

func (repository *Repository) DeleteUserById(id string) error {
	user, err := repository.FindById(id)
	if err != nil {
		return err
	}
	result := repository.db.Delete(&user, id)
	if result.RowsAffected >= 1 && result.Error != nil {
		return nil
	}
	return result.Error
}
