package handlers

import (
	"app/internal"
	"app/internal/service"
	"fmt"
	"net/http"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
)

type TicketHandler struct {
	srv service.ServiceTicketDefault
}

func NewTicketHandler(s service.ServiceTicketDefault) *TicketHandler {
	return &TicketHandler{srv: s}
}

func (e *TicketHandler) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		total, err := e.srv.GetTotalTickets(ctx)
		if err != nil {
			response.Error(w, http.StatusInternalServerError, "Erro ao buscar tickets")
			return
		}
		responseData := struct {
			Total int `json:"total"`
		}{
			Total: total,
		}
		response.JSON(w, http.StatusOK, responseData)
	}
}

func (e *TicketHandler) GetByDestination() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		country := chi.URLParam(r, "country")

		total, err := e.srv.GetTicketsByDestinationCountry(ctx, country)
		if err != nil {
			response.Error(w, http.StatusInternalServerError, "Erro ao buscar tickets")
			return
		}

		responseData := struct {
			Total map[int]internal.TicketAttributes `json:"total"`
		}{
			Total: total,
		}

		response.JSON(w, http.StatusOK, responseData)
	}
}

func (e *TicketHandler) GetAverageByDestination() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		country := chi.URLParam(r, "country")

		total, err := e.srv.GetTotalTickets(ctx)
		if err != nil {
			response.Error(w, http.StatusInternalServerError, "Erro ao buscar todos os tickets")
			return
		}

		tickets, err := e.srv.GetTicketsByDestinationCountry(ctx, country)
		if err != nil {
			response.Error(w, http.StatusInternalServerError, "Erro ao buscar tickets do país")
			return
		}

		average := float64(len(tickets)) / float64(total)
		averageStr := fmt.Sprintf(`%.1f%%`, average*100)
		responseData := struct {
			Average string `json:"average"`
		}{
			Average: averageStr,
		}

		response.JSON(w, http.StatusOK, responseData)
	}
}
