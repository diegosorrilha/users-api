package models

import (
	"fmt"
	"log"

	"github.com/diegosorrilha/users-api/db"
)

// Insert is a function to insert a user in the database.
func Insert(user User) (id int64, err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		log.Printf("Error to open connection with database: %v", err)
		return
	}
	defer conn.Close()

	sql := fmt.Sprintf("INSERT INTO users (name, age, email, password, address) VALUES ('%v', %v, '%v', '%v', '%v')",
		user.Name, user.Age, user.Email, user.Password, user.Address)

	insertResult, err := conn.Exec(sql)

	if err != nil {
		log.Printf("Error to insert user in database: %v", err)
		return
	}

	id, err = insertResult.LastInsertId()

	if err != nil {
		log.Printf("Error to get user id: %v", err)
		return
	}

	return
}
