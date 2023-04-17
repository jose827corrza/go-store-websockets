package repository

import (
	"context"

	"github.com/jose827corrza/go-store-websockets/dtos"
	"github.com/jose827corrza/go-store-websockets/models"
)

type UserRepository interface {
	//Users
	InsertUser(ctx context.Context, user *models.User) error
	GetUserById(ctx context.Context, usedId string) (*dtos.SignUpUserResponse, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	GetAllUsers(ctx context.Context) ([]*dtos.SignUpUserResponse, error)

	//Brands
	InsertBrand(ctx context.Context, brand *models.Brand) error
	GetAllUBrands(ctx context.Context) ([]*models.Brand, error)

	Close() error
}

// Repository Pattern
var implementation UserRepository

// Running time will be set
// Will be injected
func SetRepository(repository UserRepository) {
	implementation = repository
}

func InsertUser(ctx context.Context, user *models.User) error {
	return implementation.InsertUser(ctx, user)
}

func GetUserById(ctx context.Context, usedId string) (*dtos.SignUpUserResponse, error) {
	return implementation.GetUserById(ctx, usedId)
}
func GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	return implementation.GetUserByEmail(ctx, email)
}

func GetAllUsers(ctx context.Context) ([]*dtos.SignUpUserResponse, error) {
	return implementation.GetAllUsers(ctx)
}

func InsertBrand(ctx context.Context, brand *models.Brand) error {
	return implementation.InsertBrand(ctx, brand)
}

func GetAllUBrands(ctx context.Context) ([]*models.Brand, error) {
	return implementation.GetAllUBrands(ctx)
}

func Close() error {
	return implementation.Close()
}
