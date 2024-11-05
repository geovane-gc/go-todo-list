package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"todo-list/models"
	"todo-list/repositories"
	"todo-list/utils"

	"github.com/go-chi/chi"
)

func Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Error while getting ID: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	todo, err := utils.DecodeJSON[models.Todo](r)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var resp map[string]any
	rowsAffected, err := repositories.Update(int64(id), &todo)
	if err != nil {
		resp = map[string]any{"error": fmt.Sprintf("Error while updating todo: %v", err)}
	} else {
		resp = map[string]any{"message": "Todo updated successfully"}
	}

	if rowsAffected > 1 {
		log.Printf("More than one row affected while updating todo with ID: %v", id)
	}

	if err = utils.EncodeJSON(w, r, http.StatusOK, resp); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

}
