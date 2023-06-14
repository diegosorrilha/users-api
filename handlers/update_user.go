package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/diegosorrilha/users-api/models"
	"github.com/diegosorrilha/users-api/repositories"
	"github.com/go-chi/chi/v5"
)

// UpdateUser is a handler to update a user.
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	userRepo := repositories.NewMySQLUserRepository()
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		log.Printf("Error to parse id: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		log.Printf("Error to decode user: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	user.SetPassword(user.Password)
	rows, err := userRepo.Update(id, user)

	if err != nil {
		log.Printf("Error to update user with id %v: %v", id, err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	resp := map[string]string{}

	if rows > 0 {
		resp = map[string]string{
			"message": fmt.Sprintf("User updated with success. id: %v", id),
		}
	} else {
		log.Printf("No record has been updated")
		resp = map[string]string{
			"message": fmt.Sprintf("No record has been updated. id: %v", id),
		}
		w.WriteHeader(http.StatusBadRequest)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
