package repositories

import (
	"todo-list/db"
	"todo-list/models"
)

func FindOne(id int64) (todo models.Todo, err error) {
	dbConnection, err := db.OpenConnection()
	if err != nil {
		return
	}

	defer dbConnection.Close()

	err = dbConnection.QueryRow("SELECT * FROM todos WHERE id = $1", id).Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)

	return
}
