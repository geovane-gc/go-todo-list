package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"todo-list/models"
	"todo-list/repositories"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo

	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Printf("Error while decoding JSON: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var resp map[string]any
	id, err := repositories.Create(&todo)
	if err != nil {
		resp = map[string]any{"StatusCode": http.StatusInternalServerError, "Message": fmt.Sprintf("Error while creating todo: %v", err)}
	} else {
		resp = map[string]any{"StatusCode": http.StatusOK, "Message": fmt.Sprintf("Todo created successfully with ID: %v", id)}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
