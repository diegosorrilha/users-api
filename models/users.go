package models

import (
	"log"

	"github.com/diegosorrilha/users-api/crypt"
	"github.com/go-playground/validator/v10"
)

// User is the user model.
type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name" validate:"required"`
	Age      int64  `json:"age" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Address  string `json:"address"`
	Link     string `json:"link"`
}

func (user *User) Validate() error {
	validate := validator.New()

	return validate.Struct(user)
}

func (user *User) SetPassword(password string) {
	hash, err := crypt.HashPassword(user.Password)

	if err != nil {
		log.Printf("Error to try to encrypt password: %v", err)
		panic(err)
	}
	user.Password = hash

}
