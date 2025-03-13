package main

import (
	"aula/cmd/api/handlers"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	rt := chi.NewRouter()

	db := map[string]string{
		"1": "Natan",
		"2": "Lucas",
		"3": "Charles",
	}

	handler := handlers.NewEmployeeHandler(db)

	rt.Get("/employee", handler.GetAll())
	rt.Get("/employee/{id}", handler.GetByID())

	log.Println("Servidor rodando em http://localhost:8080")
	http.ListenAndServe("localhost:8080", rt)

}
