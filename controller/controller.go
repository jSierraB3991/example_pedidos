package controller

import "github.com/jsierrab3991/example_pedidos/repository"

type Controller struct {
	User     *UserController
	Article  *ArticleController
	District *DistrictController
	Client   *ClientController
	Facture  *FactureController
}

func InitiateControllers(repo *repository.Repository) *Controller {
	return &Controller{
		User:     NewUserController(repo),
		Article:  NewArticleController(repo),
		District: NewDistrictController(repo),
		Client:   NewClientController(repo),
		Facture:  NewFactureController(repo),
	}
}
