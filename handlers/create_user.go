// Package handlers collects handlers functions.
package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/diegosorrilha/users-api/models"
	"github.com/diegosorrilha/users-api/repositories"
	"github.com/diegosorrilha/users-api/responses"
)

// CreateUser is a handler to create a user.
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	userRepo := repositories.NewMySQLUserRepository()

	err := json.NewDecoder(r.Body).Decode(&user)

	var resp map[string]any

	if err != nil {
		msg := fmt.Sprintf("Error to decode body: %v", err)
		log.Print(msg)
		resp = map[string]any{
			"message": msg,
		}
		responses.FailResponse("StatusBadRequest", resp, w)
		return
	}

	user.SetPassword(user.Password)
	id, err := userRepo.Create(user)

	if err != nil {
		msg := fmt.Sprintf("Error to create user: %v", err)
		log.Print(msg)
		resp = map[string]any{
			"message": msg,
		}
		responses.FailResponse("StatusBadRequest", resp, w)

	} else {
		resp = map[string]any{
			"id":      id,
			"message": "user created with success",
		}
		responses.SuccessResponse(resp, w)
	}
}
