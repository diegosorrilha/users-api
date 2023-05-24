package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/diegosorrilha/users-api/models"
)

func ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := models.GetAll()

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadGateway)
		return

	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
