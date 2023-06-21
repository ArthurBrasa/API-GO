package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ArthurBrasa/API-GO/models"
)

func List(w http.ResponseWriter, r *http.Request) {
	todos, err := models.GetALL()
	if err != nil {
		log.Print("Erro ao obter registro: %v", err)
	}

	w.Header().Add("Content-Type", "aplication/json")
	json.NewEncoder(w).Encode(todos)

}
