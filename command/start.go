package main

import (
	"log"
	"os"

	todo "github.com/int0matar/case-book"
	"github.com/int0matar/case-book/package/handler"
	"github.com/int0matar/case-book/package/repository"
	"github.com/int0matar/case-book/package/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("failed to load environment variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		Host:     os.Getenv("DB_HOSTNAME"),
		Port:     os.Getenv("DB_PORT"),
		SSLMode:  os.Getenv("DB_MODE"),
	})

	if err != nil {
		log.Fatalf("database configuration initialization error: %s", err.Error())
	}

	repositories := repository.NewRepository(db)
	services := service.NewService(repositories)
	handlers := handler.NewHandler(services)

	server := new(todo.Server)
	if err := server.Run(os.Getenv("GIN_SERVER_PORT"), handlers.InitRoutes()); err != nil {
		log.Fatalf("an error occurred while the server was starting: %s", err.Error())
	}
}
