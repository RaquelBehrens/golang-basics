package main

import (
	"exercicio/cmd/http/router"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../../.env.example")
	if err != nil {
		log.Fatalf("Erro ao carregar arquivo .env: %v", err)
	}

	apiToken := os.Getenv("API_TOKEN")
	if apiToken == "" {
		log.Fatal("API_TOKEN não está definido no .env")
	}

	r := router.NewRouter()

	log.Println("Servidor rodando em http://localhost:8080")
	http.ListenAndServe("localhost:8080", r.MapRoutes())
}
