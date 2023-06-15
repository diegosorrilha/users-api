package repositories

import (
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestGetUserByID(t *testing.T) {
	repoSetUp()
	defer repoTearDown()

	id := 1
	userRepo := NewMySQLUserRepositoryTest()
	_, err := userRepo.GetByID(int(id))

	if err != nil {
		t.Errorf("Error to try get user with id %d: %v", id, err)

	}
}
