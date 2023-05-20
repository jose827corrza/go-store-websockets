package orm

import (
	"context"
	"log"

	"github.com/jose827corrza/go-store-websockets/models"
)

func (repo *PostgresRepository) CreateAnOrder(ctx context.Context, order *models.Order) error {

	result := repo.DB.Create(&models.Order{
		Id:         order.Id,
		Date:       order.Date,
		CustomerID: order.CustomerID,
	})
	return result.Error
}

func (repo *PostgresRepository) GetOrdersByCustomerId(ctx context.Context, customerId string) (*models.Customer, error) {
	var customerOrders *models.Customer

	result := repo.DB.Select("id", "email", "phone", "user_id").Where("id=?", customerId).Model(&models.Customer{}).Preload("Orders").First(&customerOrders).Preload("Items")
	if result.Error != nil {
		return nil, result.Error
	}
	return customerOrders, result.Error
}

func (repo *PostgresRepository) AddAProductToAnOrder(ctx context.Context, orderId string, productId string, orderItemId string) error {
	var order models.Order

	result := repo.DB.Where("id=?", orderId).Model(&models.Order{}).First(&order)
	if result.Error != nil {
		return result.Error
	}
	log.Print(order.Items)
	order.Items = append(order.Items, models.OrderItem{
		Quantity:  2, // TODO <- Set the quantities
		ProductID: productId,
		OrderID:   orderId,
		Id:        orderItemId,
	})
	log.Print("****")
	log.Print(order.Items)
	repo.DB.Save(&models.Order{
		Items:      order.Items,
		Id:         order.Id,
		CustomerID: order.CustomerID,
	})
	return result.Error
}

func (repo *PostgresRepository) GetOrderById(cxt context.Context, orderId string) (*models.Order, error) {
	var order *models.Order

	result := repo.DB.Where("id=?", orderId).Preload("Items").First(&order)
	if result.Error != nil {
		return nil, result.Error
	}
	return order, nil
}
