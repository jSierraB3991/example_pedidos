package controller

import "github.com/jsierrab3991/example_pedidos/repository"

type Controller struct {
	User    *UserController
	Article *ArticleController
}

func InitiateControllers(repo *repository.Repository) *Controller {
	return &Controller{
		User:    NewUserController(repo),
		Article: NewArticleController(repo),
	}
}
