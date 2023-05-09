package orm

import (
	"context"

	"github.com/jose827corrza/go-store-websockets/models"
)

// Receiver functions are the "methods" of the "class" alias struct
func (repo *PostgresRepository) InsertUser(ctx context.Context, user *models.User) error {
	result := repo.DB.Create(&models.User{Email: user.Email, Password: user.Password, Id: user.Id, Role: user.Role})
	return result.Error
}

func (repo *PostgresRepository) GetUserById(ctx context.Context, userId string) (*models.User, error) {
	var user = models.User{}
	// result := repo.DB.Where("id = ?", userId).First(&user)

	result := repo.DB.Select("id", "email", "role").Where("id = ?", userId).First(&user)
	return &user, result.Error
}

func (repo *PostgresRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user = models.User{}
	result := repo.DB.Where("email=?", email).First(&user)
	return &user, result.Error
}

func (repo *PostgresRepository) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	var users []*models.User
	//Thiw way retrieve all the values.
	// result := repo.DB.Find(&users)

	//Otherwise, this let us select which columns we want to obtain
	result := repo.DB.Select("id", "email", "role").Find(&users)
	return users, result.Error
}
