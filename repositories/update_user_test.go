package repositories

import (
	"testing"

	"github.com/diegosorrilha/users-api/models"
	_ "github.com/mattn/go-sqlite3"
)

func TestUpdateUser(t *testing.T) {
	repoSetUp()
	defer repoTearDown()

	id := 1
	user := models.User{
		Name:     "Jane",
		Age:      25,
		Email:    "jane@doe.com",
		Password: "123",
		Address:  "Janes street",
	}

	user.ID = int64(id)

	userRepo := NewMySQLUserRepositoryTest()
	_, err := userRepo.Update(user)

	if err != nil {
		t.Errorf("Error to try delete user with id %d: %v", id, err)
	}
}
