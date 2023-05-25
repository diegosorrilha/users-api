package models

import (
	"fmt"
	"log"

	"github.com/diegosorrilha/users-api/db"
)

// Update is a function to update a user from the database.
func Update(id int, user User) (int64, error) {
	conn, err := db.OpenConnection()

	if err != nil {
		log.Printf("Error to open connection with database: %v", err)
		return 0, err
	}
	defer conn.Close()

	res, err := conn.Exec(fmt.Sprintf("UPDATE users SET name='%v', age='%v', email='%v', password='%v', address='%v' WHERE id=%v",
		user.Name, user.Age, user.Email, user.Password, user.Address, id))

	if err != nil {
		log.Printf("Error to try update user with id %v: %v", id, err)
		return 0, nil
	}

	return res.RowsAffected()
}
