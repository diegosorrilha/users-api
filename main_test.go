package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/diegosorrilha/users-api/configs"
	"github.com/diegosorrilha/users-api/models"
	"github.com/diegosorrilha/users-api/routers"
	"github.com/go-chi/chi/v5"
)

var err_load_configs = configs.Load()

func executeRequest(req *http.Request, r *chi.Mux) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func createUser() (id any, resp *httptest.ResponseRecorder) {

	if err_load_configs == nil {
		r := routers.CreateNewRouter()

		user := models.User{
			Name:     "Jane",
			Age:      25,
			Email:    "jane@doe.com",
			Password: "123",
			Address:  "Janes street",
		}

		marshalled_user, _ := json.Marshal(user)
		req, _ := http.NewRequest("POST", "/users", bytes.NewReader(marshalled_user))

		response := executeRequest(req, r)

		response_obj := map[string]any{}

		err := json.Unmarshal(response.Body.Bytes(), &response_obj)

		if err != nil {
			return 0, nil
		}

		id := response_obj["id"]

		fmt.Printf("[SETUP] User created with success! ID: '%d' \n", id)

		return id, response

	}
	return id, resp
}

func tearDown(id any) {
	if err_load_configs == nil {

		endpoint := fmt.Sprintf("/users/%v", id)

		r := routers.CreateNewRouter()
		req, _ := http.NewRequest("DELETE", endpoint, nil)

		_ = executeRequest(req, r)

		fmt.Printf("[TEARDOWN] User deleted with success! ID: '%v' \n", id)
	}
}

// List Users tests
func TestListUsersWithSuccess(t *testing.T) {
	if err_load_configs == nil {
		r := routers.CreateNewRouter()
		req, _ := http.NewRequest("GET", "/users", nil)

		response := executeRequest(req, r)

		checkResponseCode(t, http.StatusOK, response.Code)
	}
}

func TestGetListUsersWithError404(t *testing.T) {
	if err_load_configs == nil {
		r := routers.CreateNewRouter()
		req, _ := http.NewRequest("GET", "/users/", nil)

		response := executeRequest(req, r)

		checkResponseCode(t, http.StatusNotFound, response.Code)
	}
}

// Get User tests
func TestGetUserWithSuccess(t *testing.T) {
	if err_load_configs == nil {
		id, _ := createUser()
		endpoint := fmt.Sprintf("/users/%v", id)

		r := routers.CreateNewRouter()
		req, _ := http.NewRequest("GET", endpoint, nil)

		response := executeRequest(req, r)

		checkResponseCode(t, http.StatusOK, response.Code)
	}
}

func TestGetUserNotExists(t *testing.T) {
	if err_load_configs == nil {
		r := routers.CreateNewRouter()
		req, _ := http.NewRequest("GET", "/users/999999999", nil)

		response := executeRequest(req, r)

		checkResponseCode(t, http.StatusInternalServerError, response.Code)
	}
}

func TestGetUserWithError404(t *testing.T) {
	if err_load_configs == nil {
		r := routers.CreateNewRouter()
		req, _ := http.NewRequest("GET", "/users/2/", nil)

		response := executeRequest(req, r)

		checkResponseCode(t, http.StatusNotFound, response.Code)
	}
}

func TestGetUserWithError500(t *testing.T) {
	if err_load_configs == nil {
		r := routers.CreateNewRouter()
		req, _ := http.NewRequest("GET", "/users/abc", nil)

		response := executeRequest(req, r)

		checkResponseCode(t, http.StatusInternalServerError, response.Code)
	}
}

// // Create User tests
func TestCreateUserWithSuccess(t *testing.T) {
	if err_load_configs == nil {
		id, resp := createUser()

		checkResponseCode(t, http.StatusOK, resp.Code)
		tearDown(id)

	}
}

// Delete User tests
func TestDeleteUserWithSuccess(t *testing.T) {
	if err_load_configs == nil {
		id, _ := createUser()
		endpoint := fmt.Sprintf("/users/%v", id)

		r := routers.CreateNewRouter()
		req, _ := http.NewRequest("DELETE", endpoint, nil)

		response := executeRequest(req, r)

		checkResponseCode(t, http.StatusOK, response.Code)
	}
}

func TestDeleteUserWithError400(t *testing.T) {
	if err_load_configs == nil {
		r := routers.CreateNewRouter()
		req, _ := http.NewRequest("DELETE", "/users/999999999", nil)

		response := executeRequest(req, r)

		checkResponseCode(t, http.StatusBadRequest, response.Code)
	}
}

// Update User tests
func TestUpdateUserWithSuccess(t *testing.T) {
	if err_load_configs == nil {
		id, _ := createUser()
		endpoint := fmt.Sprintf("/users/%v", id)

		r := routers.CreateNewRouter()

		user := models.User{
			Name:     "Jane Updated",
			Age:      25,
			Email:    "jane_updated@doe.com",
			Password: "123",
			Address:  "Janes street",
		}

		marshalled_user, _ := json.Marshal(user)

		req, _ := http.NewRequest("PUT", endpoint, bytes.NewReader(marshalled_user))

		response := executeRequest(req, r)

		checkResponseCode(t, http.StatusOK, response.Code)
		tearDown(id)
	}
}

func TestUpdateUserWithError(t *testing.T) {
	if err_load_configs == nil {
		r := routers.CreateNewRouter()

		user := models.User{
			Name:     "Jane Updated",
			Age:      25,
			Email:    "jane_updated@doe.com",
			Password: "123",
			Address:  "Janes street",
		}

		marshalled_user, _ := json.Marshal(user)

		req, _ := http.NewRequest("PUT", "/users/999999999", bytes.NewReader(marshalled_user))

		response := executeRequest(req, r)

		checkResponseCode(t, http.StatusBadRequest, response.Code)
	}
}
