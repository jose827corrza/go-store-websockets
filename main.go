package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/jose827corrza/go-store-websockets/handlers"
	"github.com/jose827corrza/go-store-websockets/middlewares"
	"github.com/jose827corrza/go-store-websockets/server"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error loading the env variables")
	}

	PORT := os.Getenv("PORT")
	JWT_SECRET := os.Getenv("JWT_SECRET")
	DB_URL := os.Getenv("DB_URL")

	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	DB_PSSWD := os.Getenv("DB_PSSWD")
	DB_USER := os.Getenv("DB_USER")
	DB_NAME := os.Getenv("DB_NAME")

	s, err := server.NewServer(context.Background(), &server.Config{
		Port:      PORT,
		JWTSecret: JWT_SECRET,
		DBURL:     DB_URL,
		Db_HOST:   DB_HOST,
		Db_PORT:   DB_PORT,
		Db_PSSWD:  DB_PSSWD,
		Db_USER:   DB_USER,
		Db_NAME:   DB_NAME,
	})

	if err != nil {
		log.Fatal(err)
	}

	s.Run(BindRoutes)
}

// We need to create this bind func that is requested from the method Start in package server
func BindRoutes(s server.Server, r *mux.Router) {

	// This middleware with be implemented on every path of the server
	r.Use(middlewares.CheckAuthMiddleware(s)) //TODO -> check

	r.HandleFunc("/", middlewares.LogginMiddleware(handlers.HomeHandler(s))).Methods(http.MethodGet)
	// r.HandleFunc("/", handlers.HomeHandler(s)).Methods(http.MethodGet)
	// r.HandleFunc("/signup", middlewares.LogginMiddleware(handlers.UserHandler(s))).Methods(http.MethodPost)
	r.HandleFunc("/signup", middlewares.SignUpValidator(handlers.UserHandlerInsert(s))).Methods(http.MethodPost)
	r.HandleFunc("/login", handlers.LoginHandler(s)).Methods(http.MethodPost)
	r.HandleFunc("/users/{userId}", handlers.UserHandlerGetById(s)).Methods(http.MethodGet)
	r.HandleFunc("/users", handlers.UserHandlerGetAll(s)).Methods(http.MethodGet) // <- TODO
	r.HandleFunc("/me", handlers.MeHandler(s)).Methods(http.MethodGet)
	r.HandleFunc("/brands/create", handlers.InsertNewBrand(s)).Methods(http.MethodPost)
	r.HandleFunc("/brands/{brandId}", handlers.GetAllProductsByBrandHandler(s)).Methods(http.MethodGet)
	r.HandleFunc("/products/create", handlers.InsertProductHandler(s)).Methods(http.MethodPost)
	r.HandleFunc("/products", handlers.GetAllProductsHandler(s)).Methods(http.MethodGet)
	r.HandleFunc("/products/{productId}", handlers.UpdateAProductHandler(s)).Methods(http.MethodPut)
	r.HandleFunc("/products/{productId}", handlers.DeleteAProductHandler(s)).Methods(http.MethodDelete)
	r.HandleFunc("/categories/create", handlers.InsertNewCategory(s)).Methods(http.MethodPost)
	r.HandleFunc("/categories", handlers.GetAllCategories(s)).Methods(http.MethodGet)
	r.HandleFunc("/categories/{categoryId}", handlers.GetProductsByCategory(s)).Methods(http.MethodGet)
	r.HandleFunc("/customers", handlers.CreateNewCustomerHanlder(s)).Methods(http.MethodPost)
	r.HandleFunc("/orders", handlers.CreateAnOrderHandler(s)).Methods(http.MethodPost)
	r.HandleFunc("/orders/{customerId}", handlers.GetOrdersByCustomerIdHandler(s)).Methods(http.MethodGet)
	r.HandleFunc("/orders/add-item", handlers.AddAProductToAnOrderHandler(s)).Methods(http.MethodPost)
	r.HandleFunc("/orders/items/{orderId}", handlers.GetOrderByIdHandler(s)).Methods(http.MethodGet)
}
