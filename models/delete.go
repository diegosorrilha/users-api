// Package models collects database operation functions.
package models

import (
	"fmt"
	"log"

	"github.com/diegosorrilha/users-api/db"
)

// DeleteUser is a function to delete a user in the database.
func DeleteUser(id int) (int64, error) {
	conn, err := db.OpenConnection()

	if err != nil {
		log.Printf("Error to open connection with database: %v", err)
		return 0, err
	}
	defer conn.Close()

	res, err := conn.Exec(fmt.Sprintf("DELETE from users where id=%v", id))

	if err != nil {
		log.Printf("Error to try delete user with id %v: %v", id, err)
		return 0, nil
	}

	return res.RowsAffected()
}
