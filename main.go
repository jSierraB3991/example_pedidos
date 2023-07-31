package main

import (
	"fmt"

	"github.com/jsierrab3991/example_pedidos/controller"
	"github.com/jsierrab3991/example_pedidos/database"
	"github.com/jsierrab3991/example_pedidos/libs"
	"github.com/jsierrab3991/example_pedidos/repository"
	"github.com/jsierrab3991/example_pedidos/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	c := libs.Configure("./", "postgres")
	db := database.New(c)
	database.AutoMigrate(db)
	repo := repository.InitiateRepo(db)

	controllers := controller.InitiateControllers(repo)
	routes.Routes(e, controllers)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", c.ServerPort)))
}
