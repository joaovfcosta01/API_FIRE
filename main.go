package main

import (
	"API_FIRE/configs"
	"API_FIRE/handlers"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	_ "github.com/nakagami/firebirdsql"
)

func main() {
	// Carregar configurações
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	// Criar um novo roteador chi
	r := chi.NewRouter()
	r.Post("/", handlers.Create)
	r.Put("/{id}", handlers.Update)
	r.Delete("/{id}", handlers.Delete)
	r.Get("/", handlers.List)
	r.Get("/{id}", handlers.Get)

	// Iniciar o servidor HTTP
	serverPort := configs.GetServerPort()
	addr := fmt.Sprintf(":%s", serverPort)
	fmt.Printf("Servidor rodando na porta %s\n", serverPort)
	err = http.ListenAndServe(addr, r)
	if err != nil {
		panic(err)
	}
}
