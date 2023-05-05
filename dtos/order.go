package dtos

import "github.com/jose827corrza/go-store-websockets/models"

type NewOrderRequest struct {
	CustomerID string             `validate:"required"`
	Items      []models.OrderItem `validate:"required"`
}
