package repositories

import (
	"todo-list/db"
	"todo-list/models"
)

func FindMany() (todos []models.Todo, err error) {
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
		var todo models.Todo

		err = rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)
		if err != nil {
			continue
		}

		todos = append(todos, todo)
	}

	return
}
