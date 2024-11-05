// routes/routes.go
package routes

import (
	"net/http"
	"todo-list/configs/handlers"

	"github.com/go-chi/chi"
)

func RegisterRoutes(r *chi.Mux) {
	r.Handle("/", http.NotFoundHandler())

	r.Post("/todos", handlers.Create)
	r.Get("/todos", handlers.FindMany)
	r.Get("/todos/{id}", handlers.FindOne)
	r.Put("/todos/{id}", handlers.Update)
	r.Delete("/todos/{id}", handlers.Remove)
}
