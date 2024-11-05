package routes

import (
	"net/http"
	"todo-list/configs/handlers"
	"todo-list/middlewares"

	"github.com/go-chi/chi"
)

func RegisterRoutes(r *chi.Mux) {
	r.Handle("/", http.NotFoundHandler())

	r.Post("/todos", middlewares.AdminValidation(handlers.Create))
	r.Get("/todos", http.HandlerFunc(handlers.FindMany))
	r.Get("/todos/{id}", http.HandlerFunc(handlers.FindOne))
	r.Put("/todos/{id}", middlewares.AdminValidation(handlers.Update))
	r.Delete("/todos/{id}", middlewares.AdminValidation(handlers.Remove))
}
