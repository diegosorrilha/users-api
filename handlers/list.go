package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/diegosorrilha/users-api/models"
)

func ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := models.GetAll()

	if err != nil {
		log.Printf("Error to get list of users: %v", err)
		w.WriteHeader(http.StatusBadGateway)
		return

	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
