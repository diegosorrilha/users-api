package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/diegosorrilha/users-api/models"
	"github.com/go-chi/chi/v5"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		log.Printf("Error to parse id: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user, err := models.Get(id)

	if err != nil {
		log.Printf("Error to try get user with id %v: %v", id, err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
