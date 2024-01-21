package models

import "todo-list/db"

func Insert(todo Todo) (id int64, err error) {
	dbConnection, err := db.OpenConnection()
	if err != nil {
		return
	}

	defer dbConnection.Close()

	err = dbConnection.QueryRow("INSERT INTO todos (title, description, done) VALUES ($1, $2, $3)", todo.Title, todo.Description, todo.Done).Scan(&id) // &id is a pointer to the id variable in the function return

	return
}
