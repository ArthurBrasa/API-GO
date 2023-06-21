package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/ArthurBrasa/API-GO/models"
)

func Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id")) // PARSE ID

	if err != nil {
		log.Printf("Error ao fazer parse do id: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Atuliazar
	todo, err := models.Get(int64(id))
	if err != nil {
		log.Printf("Error ao atualiar registro: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "aplication/json")
	json.NewEncoder(w).Encode(todo)
}