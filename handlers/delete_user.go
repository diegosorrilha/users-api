package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/diegosorrilha/users-api/repositories"
	"github.com/diegosorrilha/users-api/responses"
	"github.com/go-chi/chi/v5"
)

// DeleteUser is a handler to delete a user.
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	userRepo := repositories.NewMySQLUserRepository()

	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	resp := map[string]any{}

	if err != nil {
		msg := fmt.Sprintf("Error to parse id: %v", err)
		log.Print(msg)
		responses.FailResponse("InternalServerError", resp, w)

		return
	}

	rows, err := userRepo.DeleteUser(id)

	if err != nil {
		log.Printf("Error to try delete user with id %v: %v", id, err)
	}

	if rows > 0 {
		resp = map[string]any{
			"message": fmt.Sprintf("User deleted with success. id: %v", id),
		}
		responses.SuccessResponse(resp, w)
	} else {
		msg := fmt.Sprintf("No record has been deleted. id: %v", id)
		log.Print(msg)
		resp = map[string]any{
			"message": msg,
		}
		responses.FailResponse("StatusBadRequest", resp, w)
	}
}
