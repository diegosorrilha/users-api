package models

func GetAll() (users []User, err error) {

	users = []User{
		User{Name: "Diego", Age: 39, Email: "diego@sorrilha.net", Password: "****", Address: "Rua do Diego, 42"},
		User{Name: "Roberto", Age: 45, Email: "roberto@gmail.com", Password: "****", Address: "Rua do Roberto, 42"},
	}

	return
}
