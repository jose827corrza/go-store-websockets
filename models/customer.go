package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Id       string `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Phone    string `json:"phone"`
}
