package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"todo-list/configs"
	"todo-list/routes"

	"github.com/go-chi/chi"
)

func run() error {
	err := configs.Load()
	if err != nil {
		return err
	}

	r := chi.NewRouter()
	routes.RegisterRoutes(r)

	log.Println("Server started on port", configs.GetServerPort())
	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
