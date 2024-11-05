package routes

import (
	"net/http"
	"todo-list/middlewares"
	"todo-list/services"

	"github.com/go-chi/chi"
)

func RegisterRoutes(r *chi.Mux) {
	r.Handle("/", http.NotFoundHandler())

	r.Post("/todos", middlewares.AdminValidation(services.Create))
	r.Get("/todos", http.HandlerFunc(services.FindMany))
	r.Get("/todos/{id}", http.HandlerFunc(services.FindOne))
	r.Put("/todos/{id}", middlewares.AdminValidation(services.Update))
	r.Delete("/todos/{id}", middlewares.AdminValidation(services.Remove))
}
