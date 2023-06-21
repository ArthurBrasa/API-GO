package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/ArthurBrasa/API-GO/models"
	"github.com/go-chi/chi/v5"
)

func Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		log.Printf("Error ao fazer parse do id: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var todo models.Todo

	err = json.NewDecoder((r.Body)).Decode(&todo)
	if err != nil {
		log.Printf("Error ao fazer decode do json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Atuliazar
	rows, err := models.Update(int64(id), todo)
	if err != nil {
		log.Printf("Error ao atualiar registro: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if rows > 1 {
		log.Printf("Error: foram atualizados %d registro", rows)
	}

	resp := map[string]any{
		"Error":   false,
		"Message": "Dados atulizados com sucesso!",
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
