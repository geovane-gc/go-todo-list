package repositories

import (
	"todo-list/db"
	"todo-list/models"
)

func Update(id int64, todo *models.Todo) (int64, error) {
	dbConnection, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}

	defer dbConnection.Close()

	result, err := dbConnection.Exec("UPDATE todos SET title = $1, description = $2, done = $3 WHERE id = $4", todo.Title, todo.Description, todo.Done, id)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}
