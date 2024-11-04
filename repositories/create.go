package repositories

import (
	"todo-list/db"
	"todo-list/models"
)

func Create(todo *models.Todo) (id int64, err error) {
	dbConnection, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer dbConnection.Close()

	err = dbConnection.QueryRow(
		"INSERT INTO todos (title, description, done) VALUES ($1, $2, $3) RETURNING id",
		todo.Title, todo.Description, todo.Done,
	).Scan(&id)

	return
}
