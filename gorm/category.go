package orm

import (
	"context"

	"github.com/jose827corrza/go-store-websockets/models"
)

func (repo *PostgresRepository) InsertCategory(ctx context.Context, category *models.Category) error {
	result := repo.DB.Create(&models.Category{
		Id:    category.Id,
		Name:  category.Name,
		Image: category.Image,
	})

	return result.Error
}

func (repo *PostgresRepository) GetAllCategories(ctx context.Context) ([]*models.Category, error) {
	var categories []*models.Category

	result := repo.DB.Find(&categories)
	return categories, result.Error
}
