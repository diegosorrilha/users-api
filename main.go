package main

import (
	"fmt"
	"net/http"

	"github.com/diegosorrilha/users-api/configs"
	"github.com/diegosorrilha/users-api/handlers"
	"github.com/go-chi/chi/v5"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	server_path := fmt.Sprintf("localhost:%v", configs.GetServerPort())

	r := chi.NewRouter()

	// routes
	r.Get("/users", handlers.ListUsers)
	r.Get("/users/{id}", handlers.GetUser)

	fmt.Printf("Server running: http://%v/users", server_path)
	http.ListenAndServe(server_path, r)

}
