package handlers

import (
	"encoding/json"
	"errors"
	"exercicio/internal/domain"
	"fmt"
	"net/http"
	"strconv"
	"strings"

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

func (e *ProductHandler) GetConsumerPrice() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("list")
		if query == "" {
			response.Error(w, http.StatusBadRequest, "Query 'list' está vazia!")
			return
		}
		query = strings.Trim(query, "[] ")
		idStrings := strings.Split(query, ",")

		idCounts := make(map[int]int)
		for _, idStr := range idStrings {
			id, err := strconv.Atoi(strings.TrimSpace(idStr))
			if err != nil {
				response.Error(w, http.StatusBadRequest, "ID inválido na lista!")
				return
			}
			idCounts[id]++
		}

		var totalItems int = 0
		var totalPrice float64 = 0
		selectedProducts := []domain.Product{}

		for id, count := range idCounts {
			product, err := e.srv.GetByID(id)
			if err != nil {
				response.Error(w, http.StatusNotFound, "Produto não foi encontrado")
				return
			}

			if !product.IsPublished {
				errMessage := fmt.Sprintf("Produto de id %d não está publicado!", id)
				response.Error(w, http.StatusBadRequest, errMessage)
				return
			}

			if count > product.Quantity {
				response.Error(w, http.StatusBadRequest, "Quantidade pedida excede o estoque!")
				return
			}

			totalItems += count
			totalPrice += float64(count) * product.Price
			selectedProducts = append(selectedProducts, *product)
		}

		var taxRate float64
		if totalItems < 10 {
			taxRate = 0.21
		} else if totalItems <= 20 {
			taxRate = 0.17
		} else {
			taxRate = 0.15
		}

		totalPriceWithTax := totalPrice * (1 + taxRate)

		responseData := struct {
			Products   []domain.Product `json:"products"`
			TotalPrice float64          `json:"totalPrice"`
		}{
			Products:   selectedProducts,
			TotalPrice: totalPriceWithTax,
		}

		response.JSON(w, http.StatusOK, responseData)
	}
}
