package repository

import "github.com/jsierrab3991/example_pedidos/dto"

func (repo *Repository) SaveFactureOnMongo(response dto.FactureResponse) error {
	_, err := repo.collection.InsertOne(repo.ctx, response)
	return err
}
