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

func FindOne(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Error while getting ID: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var resp map[string]any
	todo, err := repositories.FindOne(int64(id))
	if err != nil {
		resp = map[string]any{"error": fmt.Sprintf("Error while getting todo: %v", err)}
	} else {
		resp = map[string]any{"data": todo}
	}

	if err = utils.EncodeJSON(w, r, http.StatusOK, resp); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
