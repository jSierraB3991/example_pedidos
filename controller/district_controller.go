package controller

import (
	"net/http"

	"github.com/devfeel/mapper"
	"github.com/jsierrab3991/example_pedidos/dto"
	"github.com/jsierrab3991/example_pedidos/models"
	"github.com/jsierrab3991/example_pedidos/repository"
	"github.com/labstack/echo/v4"
)

type DistrictController struct {
	repo *repository.Repository
}

func NewDistrictController(repo *repository.Repository) *DistrictController {
	return &DistrictController{
		repo: repo,
	}
}

func (controller *DistrictController) GetDistrict(c echo.Context) error {
	district, err := controller.repo.GetDistrict()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, controller.mapToDtos(district))
}

func (controller *DistrictController) SaveDistrict(c echo.Context) error {
	request := new(dto.DistrictDto)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	model := new(models.District)
	mapper.Mapper(request, model)
	_, err := controller.repo.SaveDistrict(model)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	response := new(dto.DistrictResponse)
	mapper.Mapper(model, response)
	return c.JSON(http.StatusCreated, response)
}

func (controller DistrictController) mapToDtos(districts []models.District) []dto.DistrictResponse {
	var dtos []dto.DistrictResponse
	for _, district := range districts {
		dto := &dto.DistrictResponse{}
		mapper.Mapper(&district, dto)
		dto.ID = district.ID
		dtos = append(dtos, *dto)
	}
	return dtos
}
