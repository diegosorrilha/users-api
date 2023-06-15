package repositories

import (
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestGetAllUsers(t *testing.T) {
	repoSetUp()
	defer repoTearDown()

	userRepo := NewMySQLUserRepositoryTest()
	_, err := userRepo.GetAll()

	if err != nil {
		t.Errorf("Error to get list of users: %s", err)
	}
}
