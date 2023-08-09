package routes

import (
	"github.com/jsierrab3991/example_pedidos/controller"
	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo, controllers *controller.Controller) {
	UserRoutes(e, controllers.User)
	ArticleRoutes(e, controllers.Article)
}

func UserRoutes(e *echo.Echo, controller *controller.UserController) {
	e.GET("/user", controller.GetUsers)
	e.GET("/user/:id", controller.GetUserById)
	e.POST("/user", controller.SaveUser)
	e.PUT("/user", controller.UpdateUser)
	e.DELETE("/user/:id", controller.DeleteUserById)
}

func ArticleRoutes(e *echo.Echo, controller *controller.ArticleController) {
	e.GET("/article", controller.GetArticles)
	e.POST("/article", controller.SaveArticle)
}
