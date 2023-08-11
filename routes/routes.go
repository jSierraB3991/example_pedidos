package routes

import (
	"github.com/jsierrab3991/example_pedidos/controller"
	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo, controllers *controller.Controller) {
	UserRoutes(e, controllers.User)
	ArticleRoutes(e, controllers.Article)
	DistrictRoutes(e, controllers.District)
	ClientRoutes(e, controllers.Client)
	FactureRoutes(e, controllers.Facture)
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

func DistrictRoutes(e *echo.Echo, controller *controller.DistrictController) {
	e.GET("/district", controller.GetDistrict)
	e.POST("/district", controller.SaveDistrict)
}

func ClientRoutes(e *echo.Echo, controller *controller.ClientController) {
	e.GET("/client", controller.GetClients)
	e.POST("/client", controller.SaveClient)
}

func FactureRoutes(e *echo.Echo, controller *controller.FactureController) {
	e.POST("/shop", controller.Shop)
	e.POST("/facture", controller.FinishShop)
}
