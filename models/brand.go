package models

import "gorm.io/gorm"

type Brand struct {
	gorm.Model
	Id       string    `json:"id,omitempty"`
	Name     string    `json:"name"`
	Image    string    `json:"image"`
	Products []Product `json:"products"`
}
