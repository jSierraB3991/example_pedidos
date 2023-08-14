package main

import (
	"context"
	"fmt"

	"github.com/devfeel/mapper"
	"github.com/jsierrab3991/example_pedidos/controller"
	"github.com/jsierrab3991/example_pedidos/database"
	"github.com/jsierrab3991/example_pedidos/dto"
	"github.com/jsierrab3991/example_pedidos/libs"
	"github.com/jsierrab3991/example_pedidos/models"
	"github.com/jsierrab3991/example_pedidos/repository"
	"github.com/jsierrab3991/example_pedidos/routes"
	"github.com/labstack/echo/v4"
)

func init() {
	mapper.Register(&models.User{})
	mapper.Register(&dto.UserDto{})

	mapper.Register(&models.Article{})
	mapper.Register(&dto.ArticleDto{})

	mapper.Register(&models.District{})
	mapper.Register(&dto.DistrictDto{})
	mapper.Register(&dto.DistrictResponse{})

	mapper.Register(&models.Client{})
	mapper.Register(&dto.ClientRequest{})
	mapper.Register(&dto.ClientResponse{})
}

func main() {
	e := echo.New()

	c := libs.Configure("./", "data")
	db := database.New(c)
	ctx := context.TODO()
	mongoCollection := database.ConnectToMongo(c.MongoUri, ctx)
	database.AutoMigrate(db)
	repo := repository.InitiateRepo(db, mongoCollection, ctx)

	controllers := controller.InitiateControllers(repo)
	routes.Routes(e, controllers)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", c.ServerPort)))
}
