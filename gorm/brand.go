package orm

import (
	"context"

	"github.com/jose827corrza/go-store-websockets/models"
)

func (repo *PostgresRepository) InsertBrand(ctx context.Context, user *models.Brand) error {
	result := repo.DB.Create(&models.Brand{Name: user.Name, Image: user.Image, Id: user.Id})
	return result.Error
}

func (repo *PostgresRepository) GetAllUBrands(ctx context.Context) ([]*models.Brand, error) {
	var brands []*models.Brand
	err := repo.DB.Model(&models.Brand{}).Preload("Products").Find(&brands).Error
	return brands, err
}
