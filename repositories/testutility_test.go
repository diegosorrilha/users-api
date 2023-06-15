package repositories

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const testDBfile string = "users.db"

// getTestDBConn is a function to open connection with test database
func getTestDBConn() (conn *sql.DB, err error) {
	conn, err = sql.Open("sqlite3", testDBfile)

	return conn, err
}

// setUp is a function to create table in test database
func repoSetUp() {
	conn, err := getTestDBConn()

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	createTableQuery := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER NOT NULL PRIMARY KEY,
		name TEXT,
		age INTEGER NOT NULL,
		email TEXT NOT NULL,
		password LONGTEXT NOT NULL,
		address TEXT NOT NULL
	);
`
	_, err = conn.Exec(createTableQuery)

	if err != nil {
		log.Fatal(err)
	}

	insertQuery := `
	INSERT INTO users (name, age, email, password, address) VALUES (
		"John Doe", 30, "john@gmail.com", "#000!", "John's street"
	);
`
	_, err = conn.Exec(insertQuery)

	if err != nil {
		log.Fatal(err)
	}
}

// tearDown is a function to drop table in test database
func repoTearDown() {
	conn, err := getTestDBConn()

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	dropTableQuery := `
		DROP TABLE IF EXISTS users;
	`
	_, err = conn.Exec(dropTableQuery)

	if err != nil {
		log.Fatal(err)
	}
}

func NewMySQLUserRepositoryTest() *MySqlUserRepository {
	conn, err := getTestDBConn()

	if err != nil {
		log.Printf("Error to open connection with database: %v", err)
		panic(err)
	}

	return &MySqlUserRepository{conn: conn}
}
