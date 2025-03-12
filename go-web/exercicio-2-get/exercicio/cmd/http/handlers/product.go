package handlers

import (
	"encoding/json"
	"errors"
	"exercicio/internal/domain"
	"fmt"
	"net/http"
	"os"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
)

type ProductHandler struct {
	srv domain.Service
}

func NewProductHandler(s domain.Service) *ProductHandler {
	return &ProductHandler{srv: s}
}

func (e *ProductHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		products, err := e.srv.GetAll()
		if err != nil {
			response.Error(w, http.StatusInternalServerError, "Erro ao buscar produtos")
			return
		}
		if len(products) == 0 {
			response.JSON(w, http.StatusNoContent, nil)
		}
		response.JSON(w, http.StatusOK, products)
	}
}

func (e *ProductHandler) GetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "productId")
		var id int
		_, err := fmt.Sscanf(idStr, "%d", &id)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "ID inválido!")
			return
		}

		products, err := e.srv.GetAll()
		if err != nil {
			response.Error(w, http.StatusInternalServerError, "Erro ao buscar produtos!")
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
			response.Error(w, http.StatusNotFound, "Produto não encontrado!")
		} else {
			response.JSON(w, http.StatusFound, result)
		}
	}
}

func (e *ProductHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token != os.Getenv("API_TOKEN") {
			response.Error(w, http.StatusUnauthorized, "Token inválido.")
			return
		}

		var reqBody domain.RequestBodyProduct
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			response.Error(w, http.StatusBadRequest, "JSON inválido!")
			return
		}

		product, err := e.srv.Create(reqBody)
		if err != nil {
			response.Error(w, http.StatusInternalServerError, err.Error())
			return
		}

		response.JSON(w, http.StatusCreated, product)
	}
}

func (e *ProductHandler) UpdateOrCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token != os.Getenv("API_TOKEN") {
			response.Error(w, http.StatusUnauthorized, "Token inválido.")
			return
		}

		idStr := chi.URLParam(r, "productId")
		var id int
		if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
			response.Error(w, http.StatusBadRequest, "ID inválido!")
			return
		}

		var reqBody domain.RequestBodyProduct
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			response.Error(w, http.StatusBadRequest, "JSON inválido!")
			return
		}

		product, err := e.srv.UpdateOrCreate(id, reqBody)
		if err != nil {
			response.Error(w, http.StatusInternalServerError, err.Error())
			return
		}

		response.JSON(w, http.StatusOK, product)
	}
}

func (e *ProductHandler) Patch() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token != os.Getenv("API_TOKEN") {
			response.Error(w, http.StatusUnauthorized, "Token inválido.")
			return
		}

		idStr := chi.URLParam(r, "productId")
		var id int
		if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
			response.Error(w, http.StatusBadRequest, "ID inválido!")
			return
		}

		var reqBody map[string]interface{}
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			response.Error(w, http.StatusBadRequest, "JSON inválido!")
			return
		}

		product, err := e.srv.Patch(id, reqBody)
		if err != nil {
			if errors.Is(err, domain.ErrResourceNotFound) {
				response.Error(w, http.StatusNotFound, "Produto não encontrado!")
				return
			}
			response.Error(w, http.StatusInternalServerError, err.Error())
			return
		}

		response.JSON(w, http.StatusOK, product)
	}
}

func (e *ProductHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token != os.Getenv("API_TOKEN") {
			response.Error(w, http.StatusUnauthorized, "Token inválido.")
			return
		}

		idStr := chi.URLParam(r, "productId")
		var id int
		if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
			response.Error(w, http.StatusBadRequest, "ID inválido!")
			return
		}

		err := e.srv.Delete(id)
		if err != nil {
			if errors.Is(err, domain.ErrResourceNotFound) {
				response.Error(w, http.StatusNotFound, "Produto não encontrado!")
				return
			}
			response.Error(w, http.StatusInternalServerError, err.Error())
			return
		}

		response.JSON(w, http.StatusNoContent, nil)
	}
}
