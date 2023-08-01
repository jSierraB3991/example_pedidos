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
	var dto dto.UserDto
	m := mapper.NewMapper()
	m.Mapper(user, dto)
	return c.JSON(http.StatusCreated, dto)
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
	return c.JSON(http.StatusCreated, controller.mapToDto(*user))
}

func (controller *UserController) UpdateUser(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	id := c.QueryParam("id")
	_, err := controller.repo.UpdateUser(user, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, controller.mapToDto(*user))
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
		dtos = append(dtos, controller.mapToDto(user))
	}
	return dtos
}

func (controller UserController) mapToDto(user models.User) dto.UserDto {
	return dto.UserDto{
		Id:       user.ID,
		Email:    user.Email,
		Password: user.Password,
	}
}
