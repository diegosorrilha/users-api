package models

import (
	"fmt"

	"github.com/diegosorrilha/users-api/configs"
	"github.com/diegosorrilha/users-api/db"
)

func GetAll() (users []User, err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return
	}
	defer conn.Close()

	rows, err := conn.Query("SELECT * from users;")
	if err != nil {
		return
	}

	for rows.Next() {
		var user User

		err = rows.Scan(&user.ID, &user.Name, &user.Age, &user.Email, &user.Password, &user.Address)
		if err != nil {
			fmt.Println(err)
			continue
		}

		user.Link = fmt.Sprintf("http://localhost:%s/users/%v", configs.GetServerPort(), user.ID)
		users = append(users, user)
	}

	return
}
