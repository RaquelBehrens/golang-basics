package main

import (
	"log"
	"net/http"
	"os"

	"exercicio/cmd/api/handlers"
	"exercicio/internal/repository"
	"exercicio/internal/service"

	"github.com/go-chi/chi/v5"
)

func main() {
	if cwd, err := os.Getwd(); err == nil {
		log.Println("Current working directory:", cwd)
	}

	repo, err := repository.NewProductRepository("../../docs/db/products.json")
	if err != nil {
		log.Fatal(err)
	}

	srv := service.NewProductService(repo)
	handler := handlers.NewProductHandler(srv)

	rt := chi.NewRouter()
	rt.Route("/products", func(rt chi.Router) {
		rt.Get("/", handler.GetAll())
		rt.Get("/{id}", handler.GetByID())
		rt.Post("/", handler.Create())
		rt.Put("/{id}", handler.UpdateOrCreate())
		rt.Patch("/{id}", handler.Patch())
		rt.Delete("/{id}", handler.Delete())
	})

	log.Println("Servidor rodando em http://localhost:8080")
	http.ListenAndServe("localhost:8080", rt)
}
