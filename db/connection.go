package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/diegosorrilha/users-api/configs"
)

func OpenConnection() (*sql.DB, error) {
	conf := configs.GetDB()

	string_connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		conf.User, conf.Pass, conf.Host, conf.Port, conf.Database)

	conn, err := sql.Open("mysql", string_connection)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	err = conn.Ping()

	return conn, err
}
