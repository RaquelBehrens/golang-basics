package handler

import (
	"app/internal"
	"app/platform/web/request"
	"app/platform/web/response"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// NewWarehousesDefault returns a new instance of WarehousesDefault
func NewWarehousesDefault(rp internal.RepositoryWarehouses) *WarehousesDefault {
	return &WarehousesDefault{
		rp: rp,
	}
}

// WarehousesDefault is a struct that represents the default warehouse handler
type WarehousesDefault struct {
	// rp is the warehouse repository
	rp internal.RepositoryWarehouses
}

// WarehouseJSON is a struct that represents a warehouse in JSON
type WarehouseJSON struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Address   string `json:"address"`
	Telephone string `json:"telephone"`
	Capacity  int    `json:"capacity"`
}

// GetOne returns a warehouse by id
func (h *WarehousesDefault) GetOne() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Error(w, http.StatusBadRequest, "invalid id")
			return
		}

		// process
		p, err := h.rp.GetOne(id)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrWarehouseNotFound):
				response.Error(w, http.StatusNotFound, "warehouse not found")
			default:
				response.Error(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		// response
		// - serialize
		data := WarehouseJSON{
			ID:        p.ID,
			Name:      p.Name,
			Address:   p.Address,
			Telephone: p.Telephone,
			Capacity:  p.Capacity,
		}
		response.JSON(w, http.StatusOK, map[string]any{"message": "warehouse found", "data": data})
	}
}

// RequestBodyWarehouseCreate is a struct that represents the request body of a warehouse to create
type RequestBodyWarehouseCreate struct {
	Name      string `json:"name"`
	Address   string `json:"address"`
	Telephone string `json:"telephone"`
	Capacity  int    `json:"capacity"`
}

// Create creates a warehouse
func (h *WarehousesDefault) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		var body RequestBodyWarehouseCreate
		if err := request.JSON(r, &body); err != nil {
			response.Error(w, http.StatusBadRequest, "invalid request body")
			return
		}

		// process
		p := internal.Warehouse{
			Name:      body.Name,
			Address:   body.Address,
			Telephone: body.Telephone,
			Capacity:  body.Capacity,
		}
		if err := h.rp.Store(&p); err != nil {
			switch {
			case errors.Is(err, internal.ErrWarehouseNotUnique):
				response.Error(w, http.StatusConflict, "warehouse not unique")
			case errors.Is(err, internal.ErrWarehouseRelation):
				response.Error(w, http.StatusConflict, "warehouse relation error")
			default:
				response.Error(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		// response
		// - serialize
		data := WarehouseJSON{
			ID:        p.ID,
			Name:      p.Name,
			Address:   p.Address,
			Telephone: p.Telephone,
			Capacity:  p.Capacity,
		}
		response.JSON(w, http.StatusCreated, map[string]any{"message": "warehouse created", "data": data})
	}
}

type WarehouseReportJSON struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Address      string `json:"address"`
	Telephone    string `json:"telephone"`
	Capacity     int    `json:"capacity"`
	ProductCount int    `json:"product_count"`
}

func (h *WarehousesDefault) GetReport() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Error(w, http.StatusBadRequest, "invalid id")
			return
		}

		warehouses, err := h.rp.ReportProducts(int(id))
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrWarehouseNotFound):
				response.Error(w, http.StatusNotFound, "warehouse not found")
			default:
				response.Error(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		var report []WarehouseReportJSON
		for _, w := range warehouses {
			data := WarehouseReportJSON{
				ID:           w.ID,
				Name:         w.Name,
				Address:      w.Address,
				Telephone:    w.Telephone,
				Capacity:     w.Capacity,
				ProductCount: w.ProductCount,
			}
			report = append(report, data)
		}

		response.JSON(w, http.StatusOK, map[string]any{"message": "report generated", "data": report})
	}
}

func (h *WarehousesDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		warehouses, err := h.rp.GetAll()
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrWarehouseNotFound):
				response.Error(w, http.StatusNotFound, "warehouse not found")
			default:
				response.Error(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		var allWarehouses []WarehouseJSON
		for _, w := range warehouses {
			data := WarehouseJSON{
				ID:        w.ID,
				Name:      w.Name,
				Address:   w.Address,
				Telephone: w.Telephone,
				Capacity:  w.Capacity,
			}
			allWarehouses = append(allWarehouses, data)
		}

		response.JSON(w, http.StatusOK, map[string]any{"message": "warehouses found", "data": allWarehouses})
	}
}
