package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Config struct {
	Port      string
	JWTSecret string
	DBURL     string
}

type Server interface {
	Config() *Config
}

type Broker struct {
	config *Config
	router *mux.Router
}

// Method for Broker that return the configuration
func (b *Broker) Config() *Config {
	return b.config
}

// Method to Create a new Broker "Constructor" method
func NewServer(ctx context.Context, config *Config) (*Broker, error) {
	if config.Port == "" {
		return nil, errors.New("port must be specified")
	}
	if config.JWTSecret == "" {
		return nil, errors.New("secret must be specified")
	}
	if config.DBURL == "" {
		return nil, errors.New("database URL must be specified")
	}
	b := &Broker{
		config: config,
		router: mux.NewRouter(),
	}
	return b, nil
}

// Method that allows to run the server
func (b *Broker) Run(binder func(s Server, r *mux.Router)) {
	b.router = mux.NewRouter()
	binder(b, b.router) // Receives b because Broker accomplishes the Server interface
	fmt.Println("Starting the server at port", b.Config().Port)
	err := http.ListenAndServe(b.config.Port, b.router)
	if err != nil {
		log.Println("Error starting the server")
	} else {
		log.Fatalf("Server stopped")
	}
}
