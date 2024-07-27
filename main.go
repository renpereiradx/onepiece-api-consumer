package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/renpereiradx/onepiece-api-consumer/handlers"
	"github.com/renpereiradx/onepiece-api-consumer/middleware"
	"github.com/renpereiradx/onepiece-api-consumer/server"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	PORT := os.Getenv("PORT")
	JWT_SECRET := os.Getenv("JWT_SECRET")
	DATABASE_URL := os.Getenv("DATABASE_URL")

	fmt.Println(DATABASE_URL)

	server, err := server.NewServer(context.Background(), &server.Config{
		Port:        PORT,
		JWTSecret:   JWT_SECRET,
		DatabaseURL: DATABASE_URL,
	})
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}

	server.Start(BindRoutes)

}

func BindRoutes(server server.Server, router *mux.Router) {
	api := router.PathPrefix("/api/v1").Subrouter()
	api.Use(middleware.CheckAuthMiddleware(server))
	router.HandleFunc("/home", handlers.HomeHandler(server)).Methods(http.MethodGet)
	router.HandleFunc("/signup", handlers.SignupHandler(server)).Methods(http.MethodPost)
}
