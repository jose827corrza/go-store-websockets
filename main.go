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

	s, err := server.NewServer(context.Background(), &server.Config{
		Port:      PORT,
		JWTSecret: JWT_SECRET,
		DBURL:     DB_URL,
	})

	if err != nil {
		log.Fatal(err)
	}

	s.Run(BindRoutes)
}

// We need to create this bind func that is requested from the method Start in package server
func BindRoutes(s server.Server, r *mux.Router) {

	// This middleware with be implemented on every path of the server
	r.Use(middlewares.CheckAuthMiddleware(s))

	r.HandleFunc("/", middlewares.LogginMiddleware(handlers.HomeHandler(s))).Methods(http.MethodGet)
	// r.HandleFunc("/", handlers.HomeHandler(s)).Methods(http.MethodGet)
	// r.HandleFunc("/signup", middlewares.LogginMiddleware(handlers.UserHandler(s))).Methods(http.MethodPost)
	r.HandleFunc("/signup", middlewares.SignUpValidator(handlers.UserHandlerInsert(s))).Methods(http.MethodPost)
	r.HandleFunc("/login", handlers.LoginHandler(s)).Methods(http.MethodPost)
	r.HandleFunc("/users/{userId}", handlers.UserHandlerGetById(s)).Methods(http.MethodGet)
	r.HandleFunc("/users", handlers.UserHandlerGetAll(s)).Methods(http.MethodGet)
	r.HandleFunc("/me", handlers.MeHandler(s)).Methods(http.MethodGet)
}
