package main

import (
	"fmt"
	"net/http"

	"github.com/diegosorrilha/users-api/configs"
	"github.com/diegosorrilha/users-api/routers"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	server_path := fmt.Sprintf("localhost:%v", configs.GetServerPort())

	r := routers.CreateNewRouter()

	fmt.Printf("Server running: http://%v/users", server_path)
	http.ListenAndServe(server_path, r)

}
