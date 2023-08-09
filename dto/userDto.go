package dto

type UserDto struct {
	ID       uint   `json:"id" mapper:"ID"`
	Email    string `json:"email" mapper:"Email"`
	Password string `json:"password" mapper:"Password"`
	Name     string `json:"name" mapper:"Name"`
	LastName string `json:"last_name" mapper:"LastName"`
	Telphone string `json:"telphone" mapper:"Telphone"`
}
