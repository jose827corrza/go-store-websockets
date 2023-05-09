package database

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/lib/pq" // Necessary

	"github.com/jose827corrza/go-store-websockets/dtos"
	"github.com/jose827corrza/go-store-websockets/models"
)

type PostgresRepository struct {
	DB *sql.DB
}

// Constructor
func NewPostgresRepository(dbUrl string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		return nil, err
	}
	return &PostgresRepository{DB: db}, nil
}

// Receiver functions are the "methods" of the "class" alias struct
func (repo *PostgresRepository) InsertUser(ctx context.Context, user *models.User) error {
	_, err := repo.DB.ExecContext(ctx, "INSERT INTO users (id,email,password) VALUES ($1, $2, $3)", user.Id, user.Email, user.Password)
	return err
}

func (repo *PostgresRepository) GetUserById(ctx context.Context, userId string) (*dtos.SignUpUserResponse, error) {
	var user = dtos.SignUpUserResponse{}
	rows, err := repo.DB.QueryContext(ctx, "SELECT id, email FROM users WHERE id=$1", userId)

	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	for rows.Next() {
		if err = rows.Scan(&user.Id, &user.Email); err == nil {
			return &user, nil
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *PostgresRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user = models.User{}
	rows, err := repo.DB.QueryContext(ctx, "SELECT id, email, password FROM users WHERE email=$1", email)

	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	for rows.Next() {
		if err = rows.Scan(&user.Id, &user.Email, &user.Password); err == nil {
			return &user, nil
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *PostgresRepository) GetAllUsers(ctx context.Context) ([]*dtos.SignUpUserResponse, error) {
	// var users []*models.User
	var users []*dtos.SignUpUserResponse

	rows, err := repo.DB.QueryContext(ctx, "SELECT id, email FROM users")
	if err != nil {
		return nil, err
	}

	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	for rows.Next() {
		// var user = models.User{}
		var user = dtos.SignUpUserResponse{}
		if err = rows.Scan(&user.Id, &user.Email); err == nil {
			users = append(users, &user)
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *PostgresRepository) Close() error {
	return repo.DB.Close()
}

func (repo *PostgresRepository) InsertBrand(ctx context.Context, user *models.Brand) error {
	_, err := repo.DB.ExecContext(ctx, "INSERT INTO brands (id,name, image) VALUES ($1, $2, $3)", user.Id, user.Name, user.Image)
	return err
}

func (repo *PostgresRepository) GetAllUBrands(ctx context.Context) ([]*models.Brand, error) {
	// var users []*models.User
	var brands []*models.Brand

	rows, err := repo.DB.QueryContext(ctx, "SELECT id, name, image FROM brands")
	if err != nil {
		return nil, err
	}

	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	for rows.Next() {
		// var user = models.User{}
		var brand = models.Brand{}
		if err = rows.Scan(&brand.Id, &brand.Name, &brand.Image); err == nil {
			brands = append(brands, &brand)
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return brands, nil
}
