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
	repo.DB.AutoMigrate(
		&models.User{},
		&models.Category{},
		&models.Brand{},
		&models.Product{},
		&models.Customer{},
	)
}

// Receiver functions are the "methods" of the "class" alias struct
func (repo *PostgresRepository) InsertUser(ctx context.Context, user *models.User) error {
	result := repo.DB.Create(&models.User{Email: user.Email, Password: user.Password, Id: user.Id})
	return result.Error
}

func (repo *PostgresRepository) GetUserById(ctx context.Context, userId string) (*dtos.SignUpUserResponse, error) {
	var user = dtos.SignUpUserResponse{}
	result := repo.DB.Where("id = ?", user.Id).First(&user)
	return &user, result.Error
}

func (repo *PostgresRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user = models.User{}
	result := repo.DB.Where("email=?", email).First(&user)
	return &user, result.Error
}

func (repo *PostgresRepository) GetAllUsers(ctx context.Context) ([]*dtos.SignUpUserResponse, error) {
	var users []*dtos.SignUpUserResponse
	result := repo.DB.Find(&users)
	return users, result.Error
}

func (repo *PostgresRepository) Close() error {
	return nil
}

func (repo *PostgresRepository) InsertBrand(ctx context.Context, user *models.Brand) error {
	result := repo.DB.Create(&models.Brand{Name: user.Name, Image: user.Image, Id: user.Id})
	return result.Error
}

func (repo *PostgresRepository) GetAllUBrands(ctx context.Context) ([]*models.Brand, error) {
	var brands []*models.Brand
	err := repo.DB.Model(&models.Brand{}).Preload("Products").Find(&brands).Error
	return brands, err
}

func (repo *PostgresRepository) GetAllProductsByBrand(ctx context.Context, brandId string) (*models.Brand, error) {
	var brand *models.Brand

	result := repo.DB.Where("id=?", brandId).Model(models.Brand{}).Preload("Products").First(&brand)
	if result.Error != nil {
		return nil, result.Error
	}
	log.Print(brand)
	return brand, nil
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
		CategoryID:  product.CategoryID,
	})
	return result.Error
}

func (repo *PostgresRepository) GetAllProducts(ctx context.Context) ([]*models.Product, error) {
	var products []*models.Product
	result := repo.DB.Find(&products)
	return products, result.Error
}

func (repo *PostgresRepository) GetAProduct(ctx context.Context, productId string) (*models.Product, error) {
	var product *models.Product
	result := repo.DB.Where("id = ?", productId).First(&product)
	return product, result.Error
}

func (repo *PostgresRepository) UpdateAProduct(ctx context.Context, productId string, product *models.Product) (*models.Product, error) {

	result := repo.DB.Where("id = ?", productId).Updates(&product).First(&product) // .First is used to be able to Return an Error of type ErrRecordNotFound
	if result.Error != nil {
		log.Print("entro al error")
		return nil, result.Error
	}
	return product, result.Error
}

func (repo *PostgresRepository) DeleteAProduct(ctx context.Context, productId string) error {
	productToDelete := repo.DB.Where("id=?", productId).First(&models.Product{})
	if productToDelete.Error != nil {
		log.Print("error does not exist")
		return productToDelete.Error
	}
	result := repo.DB.Where("id=?", productId).Delete(&models.Product{})
	if result.Error != nil {
		log.Print("error cannot be deleted")
		return productToDelete.Error
	}
	return nil
}

func (repo *PostgresRepository) InsertCustomer(ctx context.Context, customer *models.Customer) error {
	result := repo.DB.Create(&models.Customer{
		Id:       customer.Id,
		Name:     customer.Name,
		LastName: customer.LastName,
		Phone:    customer.Phone,
	})
	return result.Error
}

func (repo *PostgresRepository) InsertCategory(ctx context.Context, category *models.Category) error {
	result := repo.DB.Create(&models.Category{
		Id:    category.Id,
		Name:  category.Name,
		Image: category.Image,
	})

	return result.Error
}

func (repo *PostgresRepository) GetAllProductsByCategory(ctx context.Context, categoryId string) (*models.Category, error) {
	var category *models.Category

	result := repo.DB.Where("id=?", categoryId).Model(models.Category{}).Preload("Products").First(&category)
	if result.Error != nil {
		return nil, result.Error
	}
	log.Printf("from DB")
	log.Print(category)
	return category, nil
}

func (repo *PostgresRepository) GetAllCategories(ctx context.Context) ([]*models.Category, error) {
	var categories []*models.Category

	result := repo.DB.Find(&categories)
	return categories, result.Error
}
