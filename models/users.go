package models

import (
	"log"

	"github.com/diegosorrilha/users-api/crypt"
)

// User is the user model.
type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Age      int64  `json:"age"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Address  string `json:"address"`
	Link     string `json:"link"`
}

func (user *User) SetPassword(password string) {
	hash, err := crypt.HashPassword(user.Password)

	if err != nil {
		log.Printf("Error to try to encrypt password: %v", err)
		panic(err)
	}
	user.Password = hash

}
