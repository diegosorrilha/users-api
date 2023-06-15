package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/diegosorrilha/users-api/repositories"
	"github.com/diegosorrilha/users-api/responses"
)

// ListUsers is a handler to get all users.
func ListUsers(w http.ResponseWriter, r *http.Request) {
	userRepo := repositories.NewMySQLUserRepository()
	users, err := userRepo.GetAll()

	resp := map[string]any{}

	if err != nil {
		msg := fmt.Sprintf("Error to get list of users: %s", err)
		log.Print(msg)
		responses.FailResponse("InternalServerError", resp, w)
		return

	}

	responses.SuccessResponse(users, w)
}
