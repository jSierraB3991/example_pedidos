package repository

import (
	"errors"

	"github.com/jsierrab3991/example_pedidos/models"
	"gorm.io/gorm"
)

func (repository Repository) GetArticles() ([]models.Article, error) {
	var articles = []models.Article{}
	err := repository.db.Find(&articles).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("Table of users empty")
		}
		return nil, err
	}
	return articles, nil
}

func (repository Repository) SaveArticle(article *models.Article) (*models.Article, error) {
	result := repository.db.Create(&article)
	if result.RowsAffected >= 1 && result.Error != nil {
		return article, nil
	}
	return nil, result.Error
}
