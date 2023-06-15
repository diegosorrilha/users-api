package repositories

import (
	"testing"

	"github.com/diegosorrilha/users-api/models"
	_ "github.com/mattn/go-sqlite3"
)

func TestCreateUser(t *testing.T) {
	repoSetUp()
	defer repoTearDown()

	user := models.User{
		Name:     "Jane",
		Age:      25,
		Email:    "jane@doe.com",
		Password: "123",
		Address:  "Janes street",
	}

	userRepo := NewMySQLUserRepositoryTest()
	user.SetPassword(user.Password)

	_, err := userRepo.Create(user)

	if err != nil {
		t.Errorf("Error to create user: %v", err)

	}
}
