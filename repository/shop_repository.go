package repository

import (
	"log"

	"github.com/jsierrab3991/example_pedidos/models"
)

func (repo *Repository) SaveShop(facture *models.Facture, detail *models.DetailFacture) error {
	if facture.ID == 0 {
		result := repo.db.Create(&facture)
		if result.Error != nil {
			return result.Error
		}

		detail.FactureID = facture.ID
		return repo.saveDetail(detail)
	}
	log.Println(facture)
	return nil
}

func (repo *Repository) saveDetail(detail *models.DetailFacture) error {
	result := repo.db.Create(&detail)
	if result.RowsAffected >= 1 && result.Error != nil {
		log.Println(result.Error.Error())
		return nil
	}
	return result.Error
}
