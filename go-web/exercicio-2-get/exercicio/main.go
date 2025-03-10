package main

import (
	"encoding/json"
	"exercicio/cmd/api/handlers"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

func loadProductsFromFile(path string) ([]handlers.Product, error) {
	var products []handlers.Product

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	myDecoder := json.NewDecoder(file)
	err = myDecoder.Decode(&products)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func main() {
	products, err := loadProductsFromFile("products.json")
	if err != nil {
		panic(err)
	}
	//fmt.Println("Loaded products:", products)

	rt := chi.NewRouter()

	handler := handlers.NewProductHandler(products)

	rt.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	})

	rt.Get("/products", handler.GetAll())
	rt.Get("/products/{id}", handler.GetByID())

	log.Println("Servidor rodando em http://localhost:8080")
	http.ListenAndServe("localhost:8080", rt)

}
