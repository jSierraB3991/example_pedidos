package controller

import (
	"net/http"
	"time"

	"github.com/jsierrab3991/example_pedidos/dto"
	"github.com/jsierrab3991/example_pedidos/models"
	"github.com/jsierrab3991/example_pedidos/repository"
	"github.com/labstack/echo/v4"
)

type FactureController struct {
	repo *repository.Repository
}

func NewFactureController(repo *repository.Repository) *FactureController {
	return &FactureController{
		repo: repo,
	}
}

func (controller *FactureController) Shop(c echo.Context) error {
	request := &dto.ShopRequest{}
	err := c.Bind(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.NewError(http.StatusInternalServerError, err.Error()))
	}
	facture, detail := controller.getShopToRequest(*request)

	err = controller.repo.SaveShop(facture, detail)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.NewError(http.StatusInternalServerError, err.Error()))
	}
	return nil
}

func (controller *FactureController) FinishShop(c echo.Context) error {
	return nil
}

func (controller FactureController) getShopToRequest(request dto.ShopRequest) (*models.Facture, *models.DetailFacture) {
	return &models.Facture{
			Date:     time.Now(),
			ClientID: request.ClientID,
			UserID:   request.UserId,
			Total:    request.Total,
		},
		&models.DetailFacture{
			FactureID:    0,
			ArticleID:    request.ArticleID,
			PriceArticle: request.SubTotal,
			Stock:        request.Stock,
			Total:        request.Total,
		}
}
