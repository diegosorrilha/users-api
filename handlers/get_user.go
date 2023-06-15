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

// GetUser is a handler to get a specific user.
func GetUser(w http.ResponseWriter, r *http.Request) {
	userRepo := repositories.NewMySQLUserRepository()

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	var resp map[string]any

	if err != nil {
		msg := fmt.Sprintf("Error to parse id: %v", err)
		log.Print(msg)
		responses.FailResponse("InternalServerError", resp, w)
		return
	}

	user, err := userRepo.GetByID(id)

	if err != nil {
		msg := fmt.Sprintf("Error to try get user with id %d: %v", id, err)
		log.Print(msg)
		responses.FailResponse("InternalServerError", resp, w)
		return
	}

	responses.UserSuccessResponse(user, w)
}
