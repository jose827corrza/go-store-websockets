package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	orm "github.com/jose827corrza/go-store-websockets/gorm"
	"github.com/jose827corrza/go-store-websockets/repository"
	"github.com/jose827corrza/go-store-websockets/websockets"
)

type Config struct {
	Port      string
	JWTSecret string
	DBURL     string
	Db_HOST   string
	Db_PORT   string
	Db_PSSWD  string
	Db_USER   string
	Db_NAME   string
	SSL_MODE  string
}

type Server interface {
	Config() *Config
	Hub() *websockets.Hub
}

type Broker struct {
	config *Config
	router *mux.Router
	hub    *websockets.Hub
}

// Method for Broker that return the configuration
func (b *Broker) Config() *Config {
	return b.config
}

func (b *Broker) Hub() *websockets.Hub {
	return b.hub
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
		hub:    websockets.NewHub(),
	}
	return b, nil
}

// Method that allows to run the server
func (b *Broker) Run(binder func(s Server, r *mux.Router)) {
	b.router = mux.NewRouter()
	binder(b, b.router) // Receives b because Broker accomplishes the Server interface

	//Normal way
	// repo, err := database.NewPostgresRepository(b.Config().DBURL)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// repository.SetRepository(repo)

	//GORM way
	repo, err := orm.NewPostgresRepository(
		b.Config().Db_HOST,
		b.Config().Db_USER,
		b.Config().Db_PSSWD,
		b.Config().Db_NAME,
		b.Config().Db_PORT,
		b.Config().SSL_MODE,
	)
	if err != nil {
		log.Fatal(err)
	}

	go b.hub.Run()

	repository.SetRepository(repo)
	repo.AutoDbUpdate()
	fmt.Println("Starting the server at port", b.Config().Port)
	portFixed := fmt.Sprintf(":%s", b.config.Port)
	err = http.ListenAndServe(portFixed, b.router)
	if err != nil {
		log.Println("Error starting the server")
	} else {
		log.Fatalf("Server stopped")
	}
}
