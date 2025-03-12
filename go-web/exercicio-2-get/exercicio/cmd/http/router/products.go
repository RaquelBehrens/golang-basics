package router

import (
	"net/http"

	"exercicio/cmd/http/handlers"
	"exercicio/internal/products"

	"github.com/go-chi/chi/v5"
)

func buildProductsRoutes() http.Handler {
	rt := chi.NewRouter()

	storage := products.NewStorage("../../docs/db/products.json")
	repo := products.NewProductRepository(storage)
	srv := products.NewProductService(repo)
	handler := handlers.NewProductHandler(srv)

	pathWriteProducts := "../../docs/db/teste.json"
	allProducts, err := repo.GetAll()
	if err == nil {
		storage.WriteProducts(allProducts, pathWriteProducts)
	}

	rt.Get("/", handler.GetAll())
	rt.Get("/{productId}", handler.GetByID())
	rt.Post("/", handler.Create())
	rt.Put("/{productId}", handler.UpdateOrCreate())
	rt.Patch("/{productId}", handler.Patch())
	rt.Delete("/{productId}", handler.Delete())

	return rt
}
