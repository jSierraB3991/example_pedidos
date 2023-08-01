package routes

import (
	"github.com/jsierrab3991/example_pedidos/controller"
	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo, controllers *controller.Controller) {
	UserRoutes(e, controllers)
}

func UserRoutes(e *echo.Echo, controllers *controller.Controller) {
	userController := controllers.User
	e.GET("/user", userController.GetUsers)
	e.GET("/user/:id", userController.GetUserById)
	e.POST("/user", userController.SaveUser)
	e.PUT("/user", userController.UpdateUser)
	e.DELETE("/user/:id", userController.DeleteUserById)
}
