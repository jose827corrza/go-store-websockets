package models

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	Id        string
	Quantity  int    `json:"quantity"`
	OrderID   string `json:"orderId"`
	ProductID string `json:"productId"`
}
