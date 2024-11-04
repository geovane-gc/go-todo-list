package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"todo-list/models"
	"todo-list/repositories"

	"github.com/go-chi/chi"
)

func Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Error while getting ID: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var todo models.Todo
	err = json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Printf("Error while decoding JSON: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var resp map[string]any
	rowsAffected, err := repositories.Update(int64(id), &todo)
	if err != nil {
		resp = map[string]any{"StatusCode": http.StatusInternalServerError, "Message": fmt.Sprintf("Error while updating todo: %v", err)}
	} else {
		resp = map[string]any{"StatusCode": http.StatusOK, "Message": "Todo updated successfully"}
	}

	if rowsAffected > 1 {
		log.Printf("More than one row affected while updating todo with ID: %v", id)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

}
