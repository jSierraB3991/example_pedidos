package repository

import (
	"errors"
	"fmt"
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

func (repo *Repository) GetFacture(id string) (*[]models.DetailFacture, error) {
	var exists []models.DetailFacture
	var ids []uint
	err := repo.db.Raw("SELECT df.id FROM detail_factures df WHERE df.facture_id = ?", id).Scan(&ids)
	if err.Error != nil {
		log.Println(err.Error.Error())
		return nil, err.Error
	}

	facture := &models.Facture{}
	err = repo.db.Joins("Client").Joins("User").Find(&facture, id)
	if err.Error != nil {
		log.Println(err.Error.Error())
		return nil, err.Error
	}

	district := &models.District{}
	err = repo.db.Find(district, facture.Client.DistrictID)
	if err.Error != nil {
		log.Println(err.Error.Error())
		return nil, err.Error
	}

	if len(ids) <= 0 {
		return nil, errors.New(fmt.Sprintf("Facture Detail id: %v Not Exists", id))
	}

	repo.db.Where(ids).Joins("Article").Find(&exists)
	facture.Client.District = *district
	exists[0].Facture = *facture
	return &exists, nil
}
