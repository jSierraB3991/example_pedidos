package controller

import (
	"net/http"

	"github.com/devfeel/mapper"
	"github.com/jsierrab3991/example_pedidos/dto"
	"github.com/jsierrab3991/example_pedidos/models"
	"github.com/jsierrab3991/example_pedidos/repository"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	repo *repository.Repository
}

func NewUserController(repo *repository.Repository) *UserController {
	return &UserController{
		repo: repo,
	}
}

func (controller *UserController) GetUsers(c echo.Context) error {
	users, err := controller.repo.GetUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, controller.mapToDtos(users))
}

func (controller *UserController) GetUserById(c echo.Context) error {
	id := c.Param("id")
	user, err := controller.repo.FindById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	dto := &dto.UserDto{}
	mapper.Mapper(user, dto)
	dto.ID = user.ID

	return c.JSON(http.StatusOK, dto)
}

func (controller *UserController) SaveUser(c echo.Context) error {

	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	_, err := controller.repo.SaveUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	dto := &dto.UserDto{}
	mapper.Mapper(user, dto)
	dto.ID = user.ID
	return c.JSON(http.StatusCreated, dto)
}

func (controller *UserController) UpdateUser(c echo.Context) error {
	userDto := new(dto.UserDto)
	if err := c.Bind(userDto); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	user := new(models.User)
	mapper.Mapper(userDto, user)

	id := c.QueryParam("id")
	_, err := controller.repo.UpdateUser(user, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	dto := &dto.UserDto{}
	mapper.Mapper(user, dto)
	dto.ID = user.ID
	return c.JSON(http.StatusOK, dto)
}

func (controller *UserController) DeleteUserById(c echo.Context) error {
	id := c.Param("id")
	err := controller.repo.DeleteUserById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, nil)
}

func (controller UserController) mapToDtos(users []models.User) []dto.UserDto {
	var dtos []dto.UserDto
	for _, user := range users {
		dto := &dto.UserDto{}
		mapper.Mapper(&user, dto)
		dto.ID = user.ID
		dtos = append(dtos, *dto)
	}
	return dtos
}
