package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Id         string `json:"id"`
	Date       time.Time
	CustomerID string      // is the foreign key
	Items      []OrderItem // OrderID is the Foreign key
}
