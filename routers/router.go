// Package routers collects routes configurations.
package routers

import (
	"github.com/diegosorrilha/users-api/handlers"
	"github.com/go-chi/chi/v5"
)

// CreateNewRouter is a function to configurate routes.
func CreateNewRouter() *chi.Mux {
	r := chi.NewRouter()

	// routes
	r.Get("/users", handlers.ListUsers)
	r.Post("/users", handlers.CreateUser)
	r.Get("/users/{id}", handlers.GetUser)
	r.Delete("/users/{id}", handlers.DeleteUser)
	r.Put("/users/{id}", handlers.UpdateUser)

	return r
}
