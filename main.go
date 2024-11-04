package main

import (
	"fmt"
	"log"
	"net/http"
	"todo-list/configs"
	"todo-list/routes"

	"github.com/go-chi/chi"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	routes.RegisterRoutes(r)

	log.Println("Server started on port", configs.GetServerPort())
	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)
}
