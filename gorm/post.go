package orm

import (
	"context"

	"github.com/jose827corrza/go-store-websockets/models"
)

func (repo *PostgresRepository) CreatePostUser(ctx context.Context, user *models.User) error {
	result := repo.DB.Create(&models.UserPost{
		Email:    user.Email,
		Password: user.Password,
		Name:     user.Name,
		ID:       user.Id,
		// Model:    gorm.Model{ID: user.Model.ID},
	})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *PostgresRepository) GetUserPostById(ctx context.Context, userId string) (*models.User, error) {
	var user = models.User{}
	// result := repo.DB.Where("id = ?", userId).First(&user)

	result := repo.DB.Select("id", "email", "role").Where("id = ?", userId).First(&user)
	return &user, result.Error
}

func (repo *PostgresRepository) CreatePost(ctx context.Context, post *models.Post) error {
	result := repo.DB.Create(&models.Post{
		Title:       post.Title,
		Description: post.Description,
		ID:          post.ID,
		UserID:      post.UserID,
	})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *PostgresRepository) GetAllPosts(ctx context.Context) ([]*models.Post, error) {
	var posts []*models.Post

	result := repo.DB.Find(&posts)
	if result.Error != nil {
		return nil, result.Error
	}
	return posts, nil
}

func (repo *PostgresRepository) GetAllPostsByUserId(ctx context.Context, userId string) ([]*models.User, error) {
	var userWithPosts []*models.User

	result := repo.DB.Model(&models.User{}).Preload("Posts").Where("id=?", userId).Find(&userWithPosts)
	if result.Error != nil {
		return nil, result.Error
	}
	return userWithPosts, nil
}
