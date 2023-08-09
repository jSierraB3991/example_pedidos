package controller

import (
	"net/http"

	"github.com/devfeel/mapper"
	"github.com/jsierrab3991/example_pedidos/dto"
	"github.com/jsierrab3991/example_pedidos/models"
	"github.com/jsierrab3991/example_pedidos/repository"
	"github.com/labstack/echo/v4"
)

type ArticleController struct {
	repo *repository.Repository
}

func NewArticleController(repo *repository.Repository) *ArticleController {
	return &ArticleController{
		repo: repo,
	}
}

func (controller *ArticleController) GetArticles(c echo.Context) error {
	articles, err := controller.repo.GetArticles()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, controller.mapToDtos(articles))
}

func (controller *ArticleController) SaveArticle(c echo.Context) error {
	request := new(dto.ArticleDto)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	article := new(models.Article)
	mapper.Mapper(request, article)

	_, err := controller.repo.SaveArticle(article)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	response := &dto.ArticleDto{}
	mapper.Mapper(article, response)
	response.ID = article.ID
	return c.JSON(http.StatusCreated, response)
}

func (controller ArticleController) mapToDtos(articles []models.Article) []dto.ArticleDto {
	var dtos []dto.ArticleDto
	for _, article := range articles {
		dto := &dto.ArticleDto{}
		mapper.Mapper(&article, dto)
		dto.ID = article.ID
		dtos = append(dtos, *dto)
	}
	return dtos
}
