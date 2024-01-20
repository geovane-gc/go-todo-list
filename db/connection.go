package db

import (
	"database/sql"
	"fmt"
	"todo-list/configs"

	_ "github.com/lib/pq" // Underscore to show Go that we are going to need this package, so it won't automatically remove it
)

func OpenConnection() (*sql.DB, error) {
	conf := configs.GetDB()
	connectionInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", conf.Host, conf.Port, conf.User, conf.Pass, conf.Database)

	connection, err := sql.Open("postgres", connectionInfo)
	if err != nil {
		panic(err) // This is not a good thing to do in Production
	}

	err = connection.Ping()

	return connection, err
}
