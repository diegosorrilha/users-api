package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/diegosorrilha/users-api/models"
	"github.com/go-chi/chi/v5"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		log.Printf("Error to parse id: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	rows, err := models.DeleteUser(id)

	if err != nil {
		log.Printf("Error to try delete user with id %v: %v", id, err)
	}

	resp := map[string]string{}

	if rows > 0 {
		resp = map[string]string{
			"message": fmt.Sprintf("User deleted with success. id: %v", id),
		}
	} else {
		log.Printf("No record has been deleted")
		resp = map[string]string{
			"message": fmt.Sprintf("No record has been deleted. id: %v", id),
		}
		w.WriteHeader(http.StatusBadRequest)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
