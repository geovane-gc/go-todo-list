package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"todo-list/repositories"
)

func FindMany(w http.ResponseWriter, r *http.Request) {
	var resp map[string]any
	todos, err := repositories.FindMany()
	if err != nil {
		resp = map[string]any{"StatusCode": http.StatusInternalServerError, "Message": fmt.Sprintf("Error while listing todos: %v", err)}
	} else {
		resp = map[string]any{"StatusCode": http.StatusOK, "Data": todos}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

}
