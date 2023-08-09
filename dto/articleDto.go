package dto

type ArticleDto struct {
	ID          uint   `json:"id" mapper:"ID"`
	Name        string `json:"name" mapper:"Name"`
	Price       int    `json:"price" mapper:"Price"`
	Model       string `json:"model" mapper:"ModelArticle"`
	Description string `json:"description" mapper:"Description"`
}
