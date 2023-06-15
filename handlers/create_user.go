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

	var resp map[string]any

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		msg := fmt.Sprintf("Error to decode body: %v", err)
		log.Print(msg)
		resp = map[string]any{
			"message": msg,
		}
		responses.FailResponse("StatusBadRequest", resp, w)
		return
	}

	// validate the user struct
	err = user.Validate()

	if err != nil {
		msg := fmt.Sprintf("Error to validate payload: %v", err)
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
