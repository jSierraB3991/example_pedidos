package dto

type ShopRequest struct {
	IdShop    uint `json:"id_shop"`
	ArticleID int  `json:"article_id"`
	ClientID  int  `json:"client_id"`
	Stock     int  `json:"stock"`
	UserId    int  `json:"user_id"`
	SubTotal  int  `json:"sub_total"`
	Total     int
}
