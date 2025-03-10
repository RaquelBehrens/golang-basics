package handlers

import (
	"aula/pkg/web"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type EmployeeHandler struct {
	st map[string]string
}

func NewEmployeeHandler(db map[string]string) *EmployeeHandler {
	return &EmployeeHandler{
		st: db,
	}
}

func (e *EmployeeHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		web.ResponseJson(w, http.StatusOK, e.st, "Lista de usuários encontrada com sucesso!")
	}
}

func (e *EmployeeHandler) GetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// URL Params
		id := chi.URLParam(r, "id")
		fmt.Println("ID: ", id)

		// Query Params
		// userId := r.URL.Query().Get("userId")      // para ?userId={ID}
		// fmt.Println("userId: ", userId)

		result, exists := e.st[id]
		if !exists {
			web.ResponseJson(w, http.StatusNotFound, nil, "Usuário não encontrado!")
			return
		}

		web.ResponseJson(w, http.StatusOK, result, "Usuário encontrado com sucesso!")
	}
}
