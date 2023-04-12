package models

type Customer struct {
	Id        uint64 `json:"id"`
	Name      string `json:"name"`
	LastName  string `json:"lastName"`
	Phone     string `json:"phone"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}