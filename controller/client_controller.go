package controller

import (
	"net/http"

	"github.com/devfeel/mapper"
	"github.com/jsierrab3991/example_pedidos/dto"
	"github.com/jsierrab3991/example_pedidos/models"
	"github.com/jsierrab3991/example_pedidos/repository"
	"github.com/labstack/echo/v4"
)

type ClientController struct {
	repo *repository.Repository
}

func NewClientController(repo *repository.Repository) *ClientController {
	return &ClientController{
		repo: repo,
	}
}

func (controller *ClientController) GetClients(c echo.Context) error {
	clients, err := controller.repo.GetClients()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.NewError(http.StatusInternalServerError, err.Error()))
	}
	return c.JSON(http.StatusOK, controller.mapToDto(clients))
}

func (controller *ClientController) SaveClient(c echo.Context) error {
	request := new(dto.ClientRequest)
	err := c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.NewError(http.StatusInternalServerError, err.Error()))
	}

	client := new(models.Client)
	mapper.Mapper(request, client)

	_, err = controller.repo.SaveClient(client)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.NewError(http.StatusInternalServerError, err.Error()))
	}

	response := &dto.ClientResponse{}
	mapper.Mapper(client, response)
	return c.JSON(http.StatusCreated, response)
}

func (controller ClientController) mapToDto(clients []models.Client) []dto.ClientResponse {
	var dtos []dto.ClientResponse
	for _, client := range clients {
		dto := &dto.ClientResponse{}
		mapper.Mapper(&client, dto)
		dto.ID = client.ID
		dtos = append(dtos, *dto)
	}
	return dtos
}
