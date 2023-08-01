package dto

type UserDto struct {
	ID       uint   `json:"id" mapper:"ID"`
	Email    string `json:"email" mapper:"Email"`
	Password string `json:"password" mapper:"Password"`
}
