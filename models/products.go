package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Id          string  `json:"id,omitempty"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Image       string  `json:"image"`
	BrandID     string  `json:"brand"`
	CategoryID  string  `json:"category"`
}
