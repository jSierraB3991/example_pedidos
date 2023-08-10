package dto

type ClientResponse struct {
	ID        uint        `json:"id" mapper:"ID"`
	Name      string      `json:"name" mapper:"Name"`
	LastName  string      `json:"last_name" mapper:"LastName"`
	Telphone  string      `json:"telphone" mapper:"Telphone"`
	Email     string      `json:"email" mapper:"Email"`
	Direction string      `json:"direction" mapper:"Direction"`
	District  DistrictDto `json:"district" mapper:"District"`
}
