package repository

import (
	"context"

	"github.com/jose827corrza/go-store-websockets/models"
)

type UserRepository interface {
	InsertUser(ctx context.Context, user *models.User) error
	GetUserById(ctx context.Context, usedId string) (*models.User, error)
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

func GetUserById(ctx context.Context, usedId string) (*models.User, error) {
	return implementation.GetUserById(ctx, usedId)
}

func Close() error {
	return implementation.Close()
}
