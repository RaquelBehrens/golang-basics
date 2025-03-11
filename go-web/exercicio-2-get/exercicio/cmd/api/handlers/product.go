package handlers

import (
	"encoding/json"
	"exercicio/pkg/web"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

type Product struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Quantity    int       `json:"quantity"`
	CodeValue   string    `json:"codeValue"`
	IsPublished bool      `json:"isPublished"`
	Expiration  time.Time `json:"expiration"`
	Price       float64   `json:"price"`
}

type RequestBodyProduct struct {
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

func (e *ProductHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var reqBody RequestBodyProduct
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			web.ResponseJson(w, http.StatusBadRequest, nil, "JSON inválido!")
			return
		}

		// Verifica se os campos obrigatórios não estão vazios
		if reqBody.Name == "" || reqBody.Quantity <= 0 || reqBody.CodeValue == "" || reqBody.Expiration == "" || reqBody.Price <= 0 {
			web.ResponseJson(w, http.StatusBadRequest, nil, "Todos os campos exceto is_published são obrigatórios!")
			return
		}

		// Verifica se o code_value é único
		for _, product := range e.st {
			if product.CodeValue == reqBody.CodeValue {
				web.ResponseJson(w, http.StatusConflict, nil, "code_value já existe!")
				return
			}
		}

		// Converte a data de validade para time.Time
		expiration, err := time.Parse("02/01/2006", reqBody.Expiration)
		if err != nil {
			web.ResponseJson(w, http.StatusBadRequest, nil, "Data de validade inválida!")
			return
		}

		// Atribui um ID único
		newID := 1
		for _, product := range e.st {
			if product.ID >= newID {
				newID = product.ID + 1
			}
		}

		pr := &Product{
			ID:          newID,
			Name:        reqBody.Name,
			Quantity:    reqBody.Quantity,
			CodeValue:   reqBody.CodeValue,
			IsPublished: reqBody.IsPublished,
			Expiration:  expiration,
			Price:       reqBody.Price,
		}

		e.st = append(e.st, *pr)
		web.ResponseJson(w, http.StatusCreated, pr, "Produto adicionado!")
	}
}
