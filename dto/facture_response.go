package dto

import "time"

type DetailFactureResponse struct {
	ArticleName        string `json:"article_name" mapper:"Article.Name"`
	ArticlePrice       int    `json:"article_price" mapper:"Article.Price"`
	ArticleModel       string `json:"article_model" mapper:"Article.ModelArticle"`
	ArticleDescription string `json:"description" mapper:"Article.Description"`
	Stock              int    `json:"stock" mapper:"Stock"`
	Total              int    `json:"subtotal" mapper:"Total"`
}

type FactureResponse struct {
	IdShop         string                  `json:"id_shop"`
	Date           time.Time               `json:"date"`
	ClientName     string                  `json:"client_name"`
	ClientLastName string                  `json:"last_name"`
	ClientEmail    string                  `json:"client_email"`
	DistrictName   string                  `json:"district"`
	User           string                  `json:"user"`
	Total          int                     `json:"int"`
	Detail         []DetailFactureResponse `json:"detail"`
}
