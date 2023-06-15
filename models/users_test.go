package models

import "testing"

func TestChecksValidationValid(t *testing.T) {
	u := &User{
		Name:     "John Duo",
		Age:      45,
		Email:    "john_doe@gmail.com",
		Password: "#000!",
		Address:  "John's street ",
	}

	err := u.Validate()

	if err != nil {
		t.Error(err)
	}
}

func TestChecksValidationNotValid(t *testing.T) {
	u := &User{
		Name:    "John Duo",
		Age:     45,
		Email:   "john_doe@gmail.com",
		Address: "John's street ",
	}

	err := u.Validate()

	if err == nil {
		t.Errorf("Validation Error is expected. Got %v\n", err)
	}
}
