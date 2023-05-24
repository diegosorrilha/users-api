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

	var resp map[string]string

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resp = map[string]string{
			"message": fmt.Sprintf("Error to create user: %v", err),
		}
	} else {
		resp = map[string]string{
			"message": fmt.Sprintf("user created with success! ID: %v", id),
		}
	}

	w.Header().Add("Content-Type", "application/json")

	json.NewEncoder(w).Encode(resp)

}
