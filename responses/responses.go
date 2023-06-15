package responses

import (
	"encoding/json"
	"net/http"
)

func SetMessageResponse(resp any, w http.ResponseWriter, statusCode int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(resp)
}

func SuccessResponse(resp any, w http.ResponseWriter) {
	SetMessageResponse(resp, w, http.StatusOK)
}

func FailResponse(TypeError string, resp map[string]any, w http.ResponseWriter) {
	switch TypeError {
	case "InternalServerError":
		w.WriteHeader(http.StatusInternalServerError)

	case "StatusBadRequest":
		SetMessageResponse(resp, w, http.StatusBadRequest)
	}
}
