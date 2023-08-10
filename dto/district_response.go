package dto

type DistrictResponse struct {
	ID       uint    `json:"id" mapper:"ID"`
	Name     string  `json:"name" mapper:"Name"`
	Initials string  `json:"initials" mapper:"Initials"`
	User     UserDto `json:"user" mapper:"User"`
}
