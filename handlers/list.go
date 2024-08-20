package handlers

import (
	"API_FIRE/models"
	"encoding/json"
	"log"
	"net/http"
)

// List retorna a lista de todos os registros da tabela TODOS
func List(w http.ResponseWriter, r *http.Request) {
	// Obter todos os registros
	todos, err := models.GetAll()
	if err != nil {
		// Logar o erro e retornar um status 500 (Internal Server Error)
		log.Printf("Erro ao obter registros: %v", err)
		http.Error(w, "Erro ao obter registros", http.StatusInternalServerError)
		return
	}

	// Definir o cabeçalho Content-Type como application/json
	w.Header().Set("Content-Type", "application/json")

	// Codificar os registros como JSON e enviar a resposta
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		// Logar o erro e retornar um status 500 (Internal Server Error)
		log.Printf("Erro ao codificar registros como JSON: %v", err)
		http.Error(w, "Erro ao codificar registros como JSON", http.StatusInternalServerError)
		return
	}
}
