package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"todo-list/repositories"

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
		resp = map[string]any{"StatusCode": http.StatusInternalServerError, "Message": fmt.Sprintf("Error while getting todo: %v", err)}
	} else {
		resp = map[string]any{"StatusCode": http.StatusOK, "Data": todo}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

}
