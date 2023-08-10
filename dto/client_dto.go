package dto

type Client struct {
	Name       string `json:"name"`
	LastName   string `json:"last_name"`
	Telphone   string `json:"telphone"`
	Email      string `json:"email"`
	Direction  string `json:"direction"`
	DistrictId uint   `json:"district_id"`
}
