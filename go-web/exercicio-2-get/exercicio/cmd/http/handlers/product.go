package handlers

import (
	"encoding/json"
	"exercicio/internal/domain"
	"exercicio/pkg/web"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type ProductHandler struct {
	srv domain.ProductService
}

func NewProductHandler(s domain.ProductService) *ProductHandler {
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
			web.ResponseJson(w, http.StatusInternalServerError, nil, err.Error())
			return
		}

		web.ResponseJson(w, http.StatusCreated, product, "Produto adicionado!")
	}
}

func (e *ProductHandler) UpdateOrCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		var id int
		if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
			web.ResponseJson(w, http.StatusBadRequest, nil, "ID inválido!")
			return
		}

		var reqBody domain.RequestBodyProduct
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			web.ResponseJson(w, http.StatusBadRequest, nil, "JSON inválido!")
			return
		}

		// Mover a lógica de validação e criação para o service
		product, err := e.srv.UpdateOrCreate(id, reqBody)
		if err != nil {
			web.ResponseJson(w, http.StatusInternalServerError, nil, err.Error())
			return
		}

		web.ResponseJson(w, http.StatusOK, product, "Produto processado com sucesso!")
	}
}

func (e *ProductHandler) Patch() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		var id int
		if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
			web.ResponseJson(w, http.StatusBadRequest, nil, "ID inválido!")
			return
		}

		var reqBody map[string]interface{}
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			web.ResponseJson(w, http.StatusBadRequest, nil, "JSON inválido!")
			return
		}

		// Mover a lógica de validação e criação para o service
		product, err := e.srv.Patch(id, reqBody)
		if err != nil {
			if err.Error() == "produto não encontrado" {
				web.ResponseJson(w, http.StatusNotFound, nil, "Produto não encontrado!")
				return
			}
			web.ResponseJson(w, http.StatusInternalServerError, nil, err.Error())
			return
		}

		web.ResponseJson(w, http.StatusOK, product, "Produto atualizado com sucesso!")
	}
}

func (e *ProductHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		var id int
		if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
			web.ResponseJson(w, http.StatusBadRequest, nil, "ID inválido!")
			return
		}
		// Mover a lógica de validação e criação para o service
		err := e.srv.Delete(id)
		if err != nil {
			if err.Error() == "produto não encontrado" {
				web.ResponseJson(w, http.StatusNotFound, nil, "Produto não encontrado!")
				return
			}
			web.ResponseJson(w, http.StatusInternalServerError, nil, err.Error())
			return
		}

		web.ResponseJson(w, http.StatusOK, nil, "Produto deletado com sucesso!")
	}
}
