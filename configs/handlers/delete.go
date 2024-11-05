package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"todo-list/repositories"
	"todo-list/utils"

	"github.com/go-chi/chi"
)

func Remove(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Error while getting ID: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var resp map[string]any
	rowsAffected, err := repositories.Remove(int64(id))
	if err != nil {
		resp = map[string]any{"message": fmt.Sprintf("Error while getting todo: %v", err)}
	} else {
		resp = map[string]any{"message": fmt.Sprintf("Todo deleted successfully. Rows affected: %d", rowsAffected)}
	}

	if rowsAffected == 0 {
		resp = map[string]any{"message": fmt.Sprintf("Todo with ID %d not found", id)}
	} else if rowsAffected > 1 {
		resp = map[string]any{"message": fmt.Sprintf("Error while deleting todo with ID %d, found multiple results with same ID", id)}
	}

	if err = utils.EncodeJSON(w, r, http.StatusOK, resp); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
