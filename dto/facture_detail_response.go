package dto

type FactureDetailResponse struct {
	Article  ArticleDto `json:"article"`
	Stock    int        `json:"stock"`
	SubTotal int        `json:"sub_total"`
	Total    int
}
