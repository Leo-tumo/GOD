package main

import (
	todo "github.com/Leo-tumo/learngo/Todo-app"
	"github.com/Leo-tumo/learngo/Todo-app/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(todo.Server)
	if err := srv.Run("9898", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occurred while running http server: %s", err.Error())
	}
}
