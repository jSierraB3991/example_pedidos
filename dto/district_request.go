package dto

type DistrictDto struct {
	ID       uint   `json:"id" mapper:"ID"`
	Name     string `json:"name" mapper:"Name"`
	Initials string `json:"initials" mapper:"Initials"`
	UserID   int    `json:"user_id" mapper:"UserID"`
}
