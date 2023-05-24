package models

import (
	"fmt"
	"log"

	"github.com/diegosorrilha/users-api/db"
)

func Get(id int) (user User, err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return
	}
	defer conn.Close()

	row := conn.QueryRow(fmt.Sprintf("SELECT * from users where id=%v", id))

	err = row.Scan(&user.ID, &user.Name, &user.Age, &user.Email, &user.Password, &user.Address)
	if err != nil {
		log.Printf("Error to get user with id %v: %v", id, err)
	}

	return
}
