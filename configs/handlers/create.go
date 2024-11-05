package handlers

import (
	"fmt"
	"net/http"
	"todo-list/models"
	"todo-list/repositories"
	"todo-list/utils"
)

func Create(w http.ResponseWriter, r *http.Request) {
	todo, err := utils.DecodeJSON[models.Todo](r)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var resp map[string]any
	id, err := repositories.Create(&todo)
	if err != nil {
		resp = map[string]any{"error": fmt.Sprintf("Error while creating todo: %v", err)}
	} else {
		resp = map[string]any{"message": fmt.Sprintf("Todo created successfully with ID: %v", id)}
	}

	if err = utils.EncodeJSON(w, r, http.StatusOK, resp); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
