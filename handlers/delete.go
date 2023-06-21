package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/ArthurBrasa/API-GO/models"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		log.Printf("Error ao fazer parse do id: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Atuliazar
	rows, err := models.Delete(int64(id))
	if err != nil {
		log.Printf("Error ao remover registro: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if rows > 1 {
		log.Printf("Error: foram removidos %d registro", rows)
	}

	resp := map[string]any{
		"Error":   false,
		"Message": "Registro removido com sucesso!",
	}

	w.Header().Add("Content-Type", "aplication/json")
	json.NewEncoder(w).Encode(resp)
}
