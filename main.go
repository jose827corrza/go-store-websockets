package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/jose827corrza/go-store-websockets/handlers"
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
	r.HandleFunc("/", handlers.HomeHandler(s)).Methods(http.MethodGet)
}
