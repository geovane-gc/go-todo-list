package handlers

import (
	"fmt"
	"net/http"
	"todo-list/repositories"
	"todo-list/utils"
)

func FindMany(w http.ResponseWriter, r *http.Request) {
	var resp map[string]any
	todos, err := repositories.FindMany()
	if err != nil {
		resp = map[string]any{"message": fmt.Sprintf("Error while listing todos: %v", err)}
	} else {
		resp = map[string]any{"data": todos}
	}

	if err = utils.EncodeJSON(w, r, http.StatusOK, resp); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

}
