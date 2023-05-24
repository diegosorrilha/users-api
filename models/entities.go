package models

// name;
// age;
// email;
// password;
// address;

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Age      int64  `json:"age"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Address  string `json:"address"`
	Link     string `json:"link"`
}
