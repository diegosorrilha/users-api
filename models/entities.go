package models

// User is the user entity.
type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Age      int64  `json:"age"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Address  string `json:"address"`
	Link     string `json:"link"`
}
