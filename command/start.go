package main

import (
	"log"

	todo "github.com/int0matar/case-book"
	"github.com/int0matar/case-book/package/handler"
)

func main() {
	handler := new(handler.Handler)

	server := new(todo.Server)
	if err := server.Run("8000", handler.InitRoutes()); err != nil {
		log.Fatalf("An error occurred while the server was starting: %s", err.Error())
	}
}
