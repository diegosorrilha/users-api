package responses

import (
	"encoding/json"
	"net/http"

	"github.com/diegosorrilha/users-api/models"
)

func SetMessageResponse(resp map[string]any, w http.ResponseWriter, statusCode int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(resp)

}

func SuccessResponse(resp map[string]any, w http.ResponseWriter) {
	SetMessageResponse(resp, w, http.StatusOK)
}

func UserSuccessResponse(user models.User, w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func UsersSuccessResponse(users []models.User, w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func FailResponse(TypeError string, resp map[string]any, w http.ResponseWriter) {
	switch TypeError {
	case "InternalServerError":
		w.WriteHeader(http.StatusInternalServerError)

	case "StatusBadRequest":
		SetMessageResponse(resp, w, http.StatusBadRequest)
	}
}
