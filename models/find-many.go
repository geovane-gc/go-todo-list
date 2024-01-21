package models

import "todo-list/db"

func FindMany() (todos []Todo, err error) {
	dbConnection, err := db.OpenConnection()
	if err != nil {
		return
	}

	defer dbConnection.Close()

	rows, err := dbConnection.Query("SELECT * FROM todos")
	if err != nil {
		return
	}

	for rows.Next() {
		var todo Todo

		err = rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)
		if err != nil {
			continue
		}

		todos = append(todos, todo)
	}

	return
}
