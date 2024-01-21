package models

import "todo-list/db"

func Remove(id int64) (int64, error) {
	dbConnection, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}

	defer dbConnection.Close()

	result, err := dbConnection.Exec("DELETE FROM todos WHERE id = $1", id)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}
