package repository

import (
	"context"

	"github.com/jose827corrza/go-store-websockets/models"
)

type UserRepository interface {
	//Users
	InsertUser(ctx context.Context, user *models.User) error
	GetUserById(ctx context.Context, usedId string) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	GetAllUsers(ctx context.Context) ([]*models.User, error)

	//Brands
	InsertBrand(ctx context.Context, brand *models.Brand) error
	GetAllUBrands(ctx context.Context) ([]*models.Brand, error)
	GetAllProductsByBrand(ctx context.Context, brandId string) (*models.Brand, error)

	//Products
	InsertProduct(ctx context.Context, product *models.Product) error
	GetAllProducts(ctx context.Context) ([]*models.Product, error)
	GetAProduct(ctx context.Context, productId string) (*models.Product, error)
	UpdateAProduct(ctx context.Context, productId string, product *models.Product) (*models.Product, error)
	DeleteAProduct(ctx context.Context, productId string) error
	//Customer
	InsertCustomer(ctx context.Context, customer *models.Customer) error

	// Category
	InsertCategory(ctx context.Context, category *models.Category) error
	GetAllProductsByCategory(ctx context.Context, categoryId string) (*models.Category, error)
	GetAllCategories(ctx context.Context) ([]*models.Category, error)

	//Order
	CreateAnOrder(ctx context.Context, order *models.Order) error
	GetOrdersByCustomerId(ctx context.Context, customerId string) (*models.Customer, error)
	AddAProductToAnOrder(ctx context.Context, orderId string, productId string, orderItemId string) error
	GetOrderById(ctx context.Context, orderId string) (*models.Order, error)

	// POSTS
	CreatePostUser(ctx context.Context, user *models.User) error
	GetUserPostById(ctx context.Context, userId string) (*models.User, error)
	CreatePost(ctx context.Context, Post *models.Post) error
	GetAllPosts(ctx context.Context) ([]*models.Post, error)
	GetAllPostsByUserId(ctx context.Context, userId string) ([]*models.User, error)

	// TASKS
	CreateTask(ctx context.Context, task *models.Task) error
	GetAllTasksByUserId(ctx context.Context, userId string) ([]*models.Task, error)
	EditATaskByTaskId(ctx context.Context, taskId string, task *models.Task) (*models.Task, error)
	DeleteATaskByTaskId(ctx context.Context, taskId string) error
	CreateASubTask(ctx context.Context, subTask *models.SubTask, taskId string) (*models.Task, error)

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
func GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	return implementation.GetUserByEmail(ctx, email)
}

func GetAllUsers(ctx context.Context) ([]*models.User, error) {
	return implementation.GetAllUsers(ctx)
}

func InsertBrand(ctx context.Context, brand *models.Brand) error {
	return implementation.InsertBrand(ctx, brand)
}

func GetAllUBrands(ctx context.Context) ([]*models.Brand, error) {
	return implementation.GetAllUBrands(ctx)
}

func GetAllProductsByBrand(ctx context.Context, brandId string) (*models.Brand, error) {
	return implementation.GetAllProductsByBrand(ctx, brandId)
}

func InsertProduct(ctx context.Context, product *models.Product) error {
	return implementation.InsertProduct(ctx, product)
}

func GetAllProducts(ctx context.Context) ([]*models.Product, error) {
	return implementation.GetAllProducts(ctx)
}

func GetAProduct(ctx context.Context, productId string) (*models.Product, error) {
	return implementation.GetAProduct(ctx, productId)
}

func UpdateAProduct(ctx context.Context, productId string, product *models.Product) (*models.Product, error) {
	return implementation.UpdateAProduct(ctx, productId, product)
}

func DeleteAProduct(ctx context.Context, productId string) error {
	return implementation.DeleteAProduct(ctx, productId)
}

func InsertCustomer(ctx context.Context, customer *models.Customer) error {
	return implementation.InsertCustomer(ctx, customer)
}

func InsertCategory(ctx context.Context, category *models.Category) error {
	return implementation.InsertCategory(ctx, category)
}

func GetAllProductsByCategory(ctx context.Context, categoryId string) (*models.Category, error) {
	return implementation.GetAllProductsByCategory(ctx, categoryId)
}

func GetAllCategories(ctx context.Context) ([]*models.Category, error) {
	return implementation.GetAllCategories(ctx)
}

func CreateAnOrder(ctx context.Context, order *models.Order) error {
	return implementation.CreateAnOrder(ctx, order)
}

func GetOrdersByCustomerId(ctx context.Context, customerId string) (*models.Customer, error) {
	return implementation.GetOrdersByCustomerId(ctx, customerId)
}

func AddAProductToAnOrder(ctx context.Context, orderId string, productId string, orderItemId string) error {
	return implementation.AddAProductToAnOrder(ctx, orderId, productId, orderItemId)
}

func GetOrderById(ctx context.Context, orderId string) (*models.Order, error) {
	return implementation.GetOrderById(ctx, orderId)
}

// POSTS
func CreatePostUser(ctx context.Context, user *models.User) error {
	return implementation.CreatePostUser(ctx, user)
}

func CreatePost(ctx context.Context, post *models.Post) error {
	return implementation.CreatePost(ctx, post)
}

func GetUserPostById(ctx context.Context, userId string) (*models.User, error) {
	return implementation.GetUserPostById(ctx, userId)
}

func GetAllPosts(ctx context.Context) ([]*models.Post, error) {
	return implementation.GetAllPosts(ctx)
}

func GetAllPostsByUserId(ctx context.Context, userId string) ([]*models.User, error) {
	return implementation.GetAllPostsByUserId(ctx, userId)
}

// TASKS

func CreateTask(ctx context.Context, task *models.Task) error {
	return implementation.CreateTask(ctx, task)
}

func GetAllTasksByUserId(ctx context.Context, userId string) ([]*models.Task, error) {
	return implementation.GetAllTasksByUserId(ctx, userId)
}

func EditATaskByTaskId(ctx context.Context, taskId string, task *models.Task) (*models.Task, error) {
	return implementation.EditATaskByTaskId(ctx, taskId, task)
}

func DeleteATaskByTaskId(ctx context.Context, taskId string) error {
	return implementation.DeleteATaskByTaskId(ctx, taskId)
}

func CreateASubTask(ctx context.Context, subTask *models.SubTask, taskId string) (*models.Task, error) {
	return implementation.CreateASubTask(ctx, subTask, taskId)
}

func Close() error {
	return implementation.Close()
}
