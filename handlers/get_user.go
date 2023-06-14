package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/diegosorrilha/users-api/repositories"
	"github.com/go-chi/chi/v5"
)

// GetUser is a handler to get a specific user.
func GetUser(w http.ResponseWriter, r *http.Request) {
	userRepo := repositories.NewMySQLUserRepository()

	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		log.Printf("Error to parse id: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user, err := userRepo.Get(id)

	if err != nil {
		log.Printf("Error to try get user with id %v: %v", id, err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
