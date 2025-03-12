package router

import (
	"log"
	"net/http"

	"exercicio/cmd/http/handlers"
	"exercicio/internal/products"

	"github.com/go-chi/chi/v5"
)

func buildProductsRoutes() http.Handler {
	rt := chi.NewRouter()

	repo, err := products.NewProductRepository("../../docs/db/products.json")
	if err != nil {
		log.Fatal(err)
	}

	srv := products.NewProductService(repo)
	handler := handlers.NewProductHandler(srv)

	rt.Get("/", handler.GetAll())
	rt.Get("/{id}", handler.GetByID())
	rt.Post("/", handler.Create())
	rt.Put("/{id}", handler.UpdateOrCreate())
	rt.Patch("/{id}", handler.Patch())
	rt.Delete("/{id}", handler.Delete())

	return rt
}
