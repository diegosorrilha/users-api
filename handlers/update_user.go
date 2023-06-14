package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/diegosorrilha/users-api/models"
	"github.com/diegosorrilha/users-api/repositories"
	"github.com/diegosorrilha/users-api/responses"
	"github.com/go-chi/chi/v5"
)

// UpdateUser is a handler to update a user.
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	userRepo := repositories.NewMySQLUserRepository()
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	resp := map[string]any{}

	if err != nil {
		msg := fmt.Sprintf("Error to parse id: %v", err)
		log.Print(msg)
		responses.FailResponse("InternalServerError", resp, w)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		msg := fmt.Sprintf("Error to decode user: %v", err)
		log.Print(msg)
		responses.FailResponse("InternalServerError", resp, w)
		return
	}

	user.ID = int64(id)
	user.SetPassword(user.Password)
	rows, err := userRepo.Update(user)

	if err != nil {
		msg := fmt.Sprintf("Error to update user with id %v: %v", id, err)
		log.Print(msg)
		responses.FailResponse("InternalServerError", resp, w)
		return
	}

	if rows > 0 {
		resp = map[string]any{
			"message": fmt.Sprintf("User updated with success. id: %v", id),
		}

	} else {
		msg := fmt.Sprintf("No record has been updated. id: %v", id)
		log.Print(msg)
		resp = map[string]any{
			"message": msg,
		}
		responses.FailResponse("StatusBadRequest", resp, w)
		return

	}

	responses.SuccessResponse(resp, w)
}
