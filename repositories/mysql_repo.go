package repositories

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/diegosorrilha/users-api/configs"
	"github.com/diegosorrilha/users-api/db"
	"github.com/diegosorrilha/users-api/models"
)

type MySqlUserRepository struct {
	conn *sql.DB
}

func NewMySQLUserRepository() *MySqlUserRepository {
	conn, err := db.OpenConnection()

	if err != nil {
		log.Printf("Error to open connection with database: %v", err)
		panic(err)
	}

	return &MySqlUserRepository{conn: conn}
}

// Create is a function to create a user in the database.
func (repo *MySqlUserRepository) Create(user models.User) (id int64, err error) {
	defer repo.conn.Close()

	sql := fmt.Sprintf("INSERT INTO users (name, age, email, password, address) VALUES ('%v', %v, '%v', '%v', '%v')",
		user.Name, user.Age, user.Email, user.Password, user.Address)

	insertResult, err := repo.conn.Exec(sql)

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

// Get is a function to get a specific user from the database.
func (repo *MySqlUserRepository) Get(id int) (user models.User, err error) {
	defer repo.conn.Close()

	row := repo.conn.QueryRow(fmt.Sprintf("SELECT * from users where id=%v", id))

	err = row.Scan(&user.ID, &user.Name, &user.Age, &user.Email, &user.Password, &user.Address)
	if err != nil {
		log.Printf("Error to get user with id %v: %v", id, err)
	}

	return
}

// GetAll is a function to get all users from the database.
func (repo *MySqlUserRepository) GetAll() (users []models.User, err error) {
	defer repo.conn.Close()

	rows, err := repo.conn.Query("SELECT * from users;")
	if err != nil {
		return
	}

	for rows.Next() {
		var user models.User

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

// Update is a function to update a user from the database.
func (repo *MySqlUserRepository) Update(user models.User) (int64, error) {
	defer repo.conn.Close()

	res, err := repo.conn.Exec(fmt.Sprintf("UPDATE users SET name='%v', age='%v', email='%v', password='%v', address='%v' WHERE id=%v",
		user.Name, user.Age, user.Email, user.Password, user.Address, user.ID))

	if err != nil {
		log.Printf("Error to try update user with id %v: %v", user.ID, err)
		return 0, nil
	}

	return res.RowsAffected()
}

// DeleteUser is a function to delete a user in the database.
func (repo *MySqlUserRepository) DeleteUser(id int) (int64, error) {
	defer repo.conn.Close()

	res, err := repo.conn.Exec(fmt.Sprintf("DELETE from users where id=%v", id))

	if err != nil {
		log.Printf("Error to try delete user with id %v: %v", id, err)
		return 0, nil
	}

	return res.RowsAffected()
}
