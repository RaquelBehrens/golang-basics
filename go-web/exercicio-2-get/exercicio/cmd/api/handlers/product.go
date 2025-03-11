package handlers

import (
	"encoding/json"
	"exercicio/internal/domain"
	"exercicio/internal/service"
	"exercicio/pkg/web"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type ProductHandler struct {
	srv service.ProductService
}

func NewProductHandler(s service.ProductService) *ProductHandler {
	return &ProductHandler{srv: s}
}

func (e *ProductHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		products, err := e.srv.GetAll()
		if err != nil {
			web.ResponseJson(w, http.StatusInternalServerError, nil, "Erro ao buscar produtos")
			return
		}
		web.ResponseJson(w, http.StatusOK, products, "Lista de produtos encontrada com sucesso!")
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

		products, err := e.srv.GetAll()
		if err != nil {
			web.ResponseJson(w, http.StatusInternalServerError, nil, "Erro ao buscar produtos")
			return
		}

		var result *domain.Product
		for _, product := range products {
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

func (e *ProductHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var reqBody domain.RequestBodyProduct
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			web.ResponseJson(w, http.StatusBadRequest, nil, "JSON inválido!")
			return
		}

		// Mover a lógica de validação e criação para o service
		product, err := e.srv.Create(reqBody)
		if err != nil {
			if err.Error() == "code_value já existe" {
				web.ResponseJson(w, http.StatusConflict, nil, "code_value já existe!")
				return
			}
			if err.Error() == "dados inválidos" {
				web.ResponseJson(w, http.StatusBadRequest, nil, "Todos os campos exceto is_published são obrigatórios!")
				return
			}
			if err.Error() == "data de validade inválida" {
				web.ResponseJson(w, http.StatusBadRequest, nil, "Data de validade inválida!")
				return
			}
			web.ResponseJson(w, http.StatusInternalServerError, nil, "Erro ao criar produto")
			return
		}

		web.ResponseJson(w, http.StatusCreated, product, "Produto adicionado!")
	}
}
