package server

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/renpereiradx/onepiece-api-consumer/database/postgres"
	"github.com/renpereiradx/onepiece-api-consumer/repository"
	"github.com/rs/cors"
)

type Config struct {
	Port        string
	JWTSecret   string
	DatabaseURL string
}

type Server interface {
	Config() *Config
}

type Broker struct {
	config *Config
	router *mux.Router
}

func (b *Broker) Config() *Config {
	return b.config
}

func NewServer(ctx context.Context, config *Config) (*Broker, error) {
	if config.Port == "" {
		return nil, errors.New("port is required")
	}
	if config.JWTSecret == "" {
		return nil, errors.New("jwt secret is required")
	}
	if config.DatabaseURL == "" {
		return nil, errors.New("database url is required")
	}
	broker := &Broker{
		config: config,
		router: mux.NewRouter(),
	}
	return broker, nil
}

func (b *Broker) Start(binder func(s Server, r *mux.Router)) {
	binder(b, b.router)

	handler := cors.AllowAll().Handler(b.router)

	repo, err := postgres.NewPostgresRepository(b.config.DatabaseURL)
	if err != nil {
		log.Println("Error connecting to database")
		return
	}
	repository.SetRepository(repo)

	log.Println("Server running on port", b.config.Port)

	if err := http.ListenAndServe(b.config.Port, handler); err != nil {
		log.Println("Error starting server")
	} else {
		log.Println("Server stopped")
	}
}
