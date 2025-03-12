package main

import (
	"exercicio/cmd/http/router"
	"log"
	"net/http"
)

func main() {
	r := router.NewRouter()

	log.Println("Servidor rodando em http://localhost:8080")
	http.ListenAndServe("localhost:8080", r.MapRoutes())
}
