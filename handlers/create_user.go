package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/diegosorrilha/users-api/models"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		log.Printf("Error to decode body: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	id, err := models.Insert(user)

	var resp map[string]any

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resp = map[string]any{
			"message": fmt.Sprintf("Error to create user: %v", err),
		}
	} else {
		resp = map[string]any{
			"id":      id,
			"message": "user created with success",
		}
	}

	w.Header().Add("Content-Type", "application/json")

	json.NewEncoder(w).Encode(resp)

}
