package main

import (
	"github.com/int0matar/book-of-cases"
	"github.com/int0matar/book-of-cases/package/handler"
	"log"
)

func main() {
	handler := new(handler.Handler)

	server := new(todo.Server)
	if err := server.Run("8000", handler.InitRoutes()); err != nil {
		log.Fatalf("An error occurred while the server was starting: %s", err.Error())
	}
}
