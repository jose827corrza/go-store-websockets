package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Id       string    `json:"id,omitempty"`
	Name     string    `json:"name"`
	Products []Product `json:"products"`
}
