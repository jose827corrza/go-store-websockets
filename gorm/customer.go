package orm

import (
	"context"

	"github.com/jose827corrza/go-store-websockets/models"
)

func (repo *PostgresRepository) InsertCustomer(ctx context.Context, customer *models.Customer) error {
	result := repo.DB.Create(&models.Customer{
		Id:       customer.Id,
		Name:     customer.Name,
		LastName: customer.LastName,
		Phone:    customer.Phone,
		User: models.User{ // Supposed to create a record in users table as well
			Id:       customer.User.Id,
			Email:    customer.User.Email,
			Password: customer.User.Password,
			Role:     customer.User.Role,
		},
	})
	return result.Error
}
