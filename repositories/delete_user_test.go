package repositories

import (
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestDeleteUser(t *testing.T) {
	repoSetUp()
	defer repoTearDown()

	id := 1
	userRepo := NewMySQLUserRepositoryTest()
	_, err := userRepo.DeleteUser(int(id))

	if err != nil {
		t.Errorf("Error to try delete user with id %d: %v", id, err)
	}

}
