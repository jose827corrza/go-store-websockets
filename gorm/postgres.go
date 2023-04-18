package orm

import (
	"context"
	"fmt"
	"log"

	"github.com/jose827corrza/go-store-websockets/dtos"
	"github.com/jose827corrza/go-store-websockets/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresRepository struct {
	DB *gorm.DB
}

func NewPostgresRepository(host string, user string, password string, dbName string, port string, sslMode string) (*PostgresRepository, error) {
	// dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, user, password, dbName, port, sslMode)
	dns := "host=localhost user=joseDev password=postgres dbname=go_estore port=5432 sslmode=disable"
	fmt.Println(dns)
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to the DB")
		return nil, err
	}
	return &PostgresRepository{DB: db}, nil
}

//AutoMigrate
func (repo *PostgresRepository) AutoDbUpdate() {
	repo.DB.AutoMigrate(&models.User{}, &models.Brand{}, &models.Product{})
}

// Receiver functions are the "methods" of the "class" alias struct
func (repo *PostgresRepository) InsertUser(ctx context.Context, user *models.User) error {
	// _, err := repo.DB.ExecContext(ctx, "INSERT INTO users (id,email,password) VALUES ($1, $2, $3)", user.Id, user.Email, user.Password)
	result := repo.DB.Create(&models.User{Email: user.Email, Password: user.Password, Id: user.Id})
	return result.Error
}

func (repo *PostgresRepository) GetUserById(ctx context.Context, userId string) (*dtos.SignUpUserResponse, error) {
	var user = dtos.SignUpUserResponse{}
	// rows, err := repo.DB.QueryContext(ctx, "SELECT id, email FROM users WHERE id=$1", userId)

	// defer func() {
	// 	err = rows.Close()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }()

	// for rows.Next() {
	// 	if err = rows.Scan(&user.Id, &user.Email); err == nil {
	// 		return &user, nil
	// 	}
	// }
	// if err = rows.Err(); err != nil {
	// 	return nil, err
	// }
	// return &user, nil
	result := repo.DB.Where("id = ?", user.Id).First(&user)
	return &user, result.Error
}

func (repo *PostgresRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user = models.User{}
	// rows, err := repo.DB.QueryContext(ctx, "SELECT id, email, password FROM users WHERE email=$1", email)

	// defer func() {
	// 	err = rows.Close()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }()

	// for rows.Next() {
	// 	if err = rows.Scan(&user.Id, &user.Email, &user.Password); err == nil {
	// 		return &user, nil
	// 	}
	// }
	// if err = rows.Err(); err != nil {
	// 	return nil, err
	// }
	// return &user, nil
	result := repo.DB.Where("email=?", email).First(&user)
	return &user, result.Error
}

func (repo *PostgresRepository) GetAllUsers(ctx context.Context) ([]*dtos.SignUpUserResponse, error) {
	// var users []*models.User
	var users []*dtos.SignUpUserResponse

	// rows, err := repo.DB.QueryContext(ctx, "SELECT id, email FROM users")
	// if err != nil {
	// 	return nil, err
	// }

	// defer func() {
	// 	err = rows.Close()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }()

	// for rows.Next() {
	// 	// var user = models.User{}
	// 	var user = dtos.SignUpUserResponse{}
	// 	if err = rows.Scan(&user.Id, &user.Email); err == nil {
	// 		users = append(users, &user)
	// 	}
	// }
	// if err = rows.Err(); err != nil {
	// 	return nil, err
	// }
	// return users, nil
	result := repo.DB.Find(&users)
	return users, result.Error
}

func (repo *PostgresRepository) Close() error {
	return nil
}

func (repo *PostgresRepository) InsertBrand(ctx context.Context, user *models.Brand) error {
	// _, err := repo.DB.ExecContext(ctx, "INSERT INTO brands (id,name, image) VALUES ($1, $2, $3)", user.Id, user.Name, user.Image)
	result := repo.DB.Create(&models.Brand{Name: user.Name, Image: user.Image, Id: user.Id})
	return result.Error
}

func (repo *PostgresRepository) GetAllUBrands(ctx context.Context) ([]*models.Brand, error) {
	// var users []*models.User
	var brands []*models.Brand

	// rows, err := repo.DB.QueryContext(ctx, "SELECT id, name, image FROM brands")
	// if err != nil {
	// 	return nil, err
	// }

	// defer func() {
	// 	err = rows.Close()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }()

	// for rows.Next() {
	// 	// var user = models.User{}
	// 	var brand = models.Brand{}
	// 	if err = rows.Scan(&brand.Id, &brand.Name, &brand.Image); err == nil {
	// 		brands = append(brands, &brand)
	// 	}
	// }
	// if err = rows.Err(); err != nil {
	// 	return nil, err
	// }
	// return brands, nil
	// result := repo.DB.Find(&brands)
	err := repo.DB.Model(&models.Brand{}).Preload("Products").Find(&brands).Error
	return brands, err
}

func (repo *PostgresRepository) InsertProduct(ctx context.Context, product *models.Product) error {
	result := repo.DB.Create(&models.Product{
		Id:          product.Id,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
		Image:       product.Image,
		BrandID:     product.BrandID,
	})
	return result.Error
}
