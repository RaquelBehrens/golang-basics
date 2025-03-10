package handlers

import (
	"exercicio/pkg/web"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"codeValue"`
	IsPublished bool    `json:"isPublished"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

type ProductHandler struct {
	st []Product
}

func NewProductHandler(db []Product) *ProductHandler {
	return &ProductHandler{
		st: db,
	}
}

func (e *ProductHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		web.ResponseJson(w, http.StatusOK, e.st, "Lista de usuários encontrada com sucesso!")
	}
}

func (e *ProductHandler) GetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		var id int
		_, err := fmt.Sscanf(idStr, "%d", &id) // Coleta o ID da URL e converte para int
		if err != nil {
			web.ResponseJson(w, http.StatusBadRequest, nil, "ID inválido!")
			return
		}

		var result *Product
		for _, product := range e.st {
			if product.ID == id {
				result = &product
				break
			}
		}

		if result == nil {
			web.ResponseJson(w, http.StatusNotFound, nil, "Produto não encontrado!")
		} else {
			web.ResponseJson(w, http.StatusOK, result, "Produto encontrado com sucesso!")
		}
	}
}
