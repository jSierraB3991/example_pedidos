package main

import (
	"net/http"

	"github.com/jsierrab3991/example_pedidos/database"
	"github.com/jsierrab3991/example_pedidos/libs"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	c := libs.Configure("./", "postgres")
	db := database.New(c)
	database.AutoMigrate(db)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":3000"))
}
