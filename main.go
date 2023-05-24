package main

import (
	"fmt"
	"net/http"

	"github.com/diegosorrilha/users-api/configs"
	"github.com/diegosorrilha/users-api/handlers"
)

func main() {
	port := 8000
	server_address := "localhost"
	server_path := fmt.Sprintf("%v:%v", server_address, port)

	err := configs.Load()
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()

	// routes
	mux.HandleFunc("/users", handlers.ListUsers)

	fmt.Printf("Server running: http://%v/users", server_path)
	http.ListenAndServe(server_path, mux)

}
