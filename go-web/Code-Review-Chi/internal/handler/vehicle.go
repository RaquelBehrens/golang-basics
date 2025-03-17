package handler

import (
	"app/internal"
	"app/internal/domain"
	"app/internal/loader"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
)

// VehicleJSON is a struct that represents a vehicle in JSON format
type VehicleJSON struct {
	ID              int     `json:"id"`
	Brand           string  `json:"brand"`
	Model           string  `json:"model"`
	Registration    string  `json:"registration"`
	Color           string  `json:"color"`
	FabricationYear int     `json:"year"`
	Capacity        int     `json:"passengers"`
	MaxSpeed        float64 `json:"max_speed"`
	FuelType        string  `json:"fuel_type"`
	Transmission    string  `json:"transmission"`
	Weight          float64 `json:"weight"`
	Height          float64 `json:"height"`
	Length          float64 `json:"length"`
	Width           float64 `json:"width"`
}

// NewVehicleDefault is a function that returns a new instance of VehicleDefault
func NewVehicleDefault(sv internal.VehicleService) *VehicleDefault {
	return &VehicleDefault{sv: sv}
}

// VehicleDefault is a struct with methods that represent handlers for vehicles
type VehicleDefault struct {
	// sv is the service that will be used by the handler
	sv internal.VehicleService
}

// GetAll is a method that returns a handler for the route GET /vehicles
func (h *VehicleDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...

		// process
		// - get all vehicles
		v, err := h.sv.FindAll()
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

		// response
		data := make(map[int]VehicleJSON)
		for key, value := range v {
			data[key] = VehicleJSON{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:        value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

func (h *VehicleDefault) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		var reqBody loader.VehicleJSON
		err := json.NewDecoder(r.Body).Decode(&reqBody)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, "Dados do veículo mal formatados ou incompletos.")
			return
		}

		vehicle := internal.Vehicle{
			Id: reqBody.Id,
			VehicleAttributes: internal.VehicleAttributes{
				Brand:           reqBody.Brand,
				Model:           reqBody.Model,
				Registration:    reqBody.Registration,
				Color:           reqBody.Color,
				FabricationYear: reqBody.FabricationYear,
				Capacity:        reqBody.Capacity,
				MaxSpeed:        reqBody.MaxSpeed,
				FuelType:        reqBody.FuelType,
				Transmission:    reqBody.Transmission,
				Weight:          reqBody.Weight,
				Dimensions: internal.Dimensions{
					Height: reqBody.Height,
					Length: reqBody.Length,
					Width:  reqBody.Width,
				},
			},
		}

		// process
		// create vehicle
		err = h.sv.Create(vehicle)
		if err != nil {
			if errors.Is(err, domain.ErrIdAlreadyExists) {
				response.JSON(w, http.StatusConflict, "Identificador do veículo já existente.")
				return
			} else {
				response.JSON(w, http.StatusInternalServerError, nil)
				return
			}
		}
		response.JSON(w, http.StatusCreated, "Veículo criado com sucesso.")
	}
}

func (h *VehicleDefault) GetByColorYear() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		color := chi.URLParam(r, "color")

		yearStr := chi.URLParam(r, "year")
		var year int
		if _, err := fmt.Sscanf(yearStr, "%d", &year); err != nil {
			response.Error(w, http.StatusBadRequest, "Ano inválido")
		}

		// process
		// - get all vehicles
		v, err := h.sv.GetByColorYear(color, year)
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

		if len(v) <= 0 {
			response.JSON(w, http.StatusNotFound, "Nenhum veículo encontrado com esses critérios.")
			return
		}

		// response
		data := make(map[int]VehicleJSON)
		for key, value := range v {
			data[key] = VehicleJSON{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:        value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}
		response.JSON(w, http.StatusOK, data)
	}
}

func (h *VehicleDefault) GetByBrandFabricatedBetween() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		brand := chi.URLParam(r, "brand")
		startYearStr := chi.URLParam(r, "start_year")
		endYearStr := chi.URLParam(r, "end_year")
		var sy, ey int

		fmt.Sscanf(startYearStr, "%d", &sy)
		fmt.Sscanf(endYearStr, "%d", &ey)

		// process
		// - get all vehicles
		v, err := h.sv.GetByBrandFabricatedBetween(brand, sy, ey)
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

		if len(v) <= 0 {
			response.JSON(w, http.StatusNotFound, "Nenhum veículo encontrado com esses critérios")
			return
		}

		// response
		data := make(map[int]VehicleJSON)
		for key, value := range v {
			data[key] = VehicleJSON{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:        value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}
		response.JSON(w, http.StatusOK, data)
	}
}

func (h *VehicleDefault) GetAverageSpeedByBrand() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		brand := chi.URLParam(r, "brand")

		// process
		// - get all vehicles
		as, err := h.sv.GetAverageSpeedByBrand(brand)

		if err != nil {
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

		if as <= 0 {
			response.JSON(w, http.StatusNotFound, "Nenhum veículo encontrado dessa marca")
			return
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message":       "success",
			"average_speed": as,
		})
	}
}

func (h *VehicleDefault) CreateBatch() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		var reqBody []loader.VehicleJSON
		err := json.NewDecoder(r.Body).Decode(&reqBody)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, "Dados de algum veículo mal formatados ou incompletos")
		}

		var vehicles []internal.Vehicle

		for _, reqVehicle := range reqBody {
			vehicle := internal.Vehicle{
				Id: reqVehicle.Id,
				VehicleAttributes: internal.VehicleAttributes{
					Brand:           reqVehicle.Brand,
					Model:           reqVehicle.Model,
					Registration:    reqVehicle.Registration,
					Color:           reqVehicle.Color,
					FabricationYear: reqVehicle.FabricationYear,
					Capacity:        reqVehicle.Capacity,
					MaxSpeed:        reqVehicle.MaxSpeed,
					FuelType:        reqVehicle.FuelType,
					Transmission:    reqVehicle.Transmission,
					Weight:          reqVehicle.Weight,
					Dimensions: internal.Dimensions{
						Height: reqVehicle.Height,
						Length: reqVehicle.Length,
						Width:  reqVehicle.Width,
					},
				},
			}
			vehicles = append(vehicles, vehicle)
		}

		// process
		err = h.sv.CreateBatch(vehicles)
		if err != nil {
			if errors.Is(err, domain.ErrIdAlreadyExists) {
				response.JSON(w, http.StatusConflict, "Algum veículo possui um identificador já existente.")
				return
			}
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

		response.JSON(w, http.StatusCreated, "Veículos criados com sucesso.")
	}
}

func (h *VehicleDefault) UpdateSpeed() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		idStr := chi.URLParam(r, "id")
		var id int
		fmt.Sscanf(idStr, "%d", &id)

		var reqBody loader.VehicleJSON
		err := json.NewDecoder(r.Body).Decode(&reqBody)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, "Velocidade mal formatada ou fora de alcance.")
			return
		}

		vehicle := internal.Vehicle{
			Id: reqBody.Id,
			VehicleAttributes: internal.VehicleAttributes{
				Brand:           reqBody.Brand,
				Model:           reqBody.Model,
				Registration:    reqBody.Registration,
				Color:           reqBody.Color,
				FabricationYear: reqBody.FabricationYear,
				Capacity:        reqBody.Capacity,
				MaxSpeed:        reqBody.MaxSpeed,
				FuelType:        reqBody.FuelType,
				Transmission:    reqBody.Transmission,
				Weight:          reqBody.Weight,
				Dimensions: internal.Dimensions{
					Height: reqBody.Height,
					Length: reqBody.Length,
					Width:  reqBody.Width,
				},
			},
		}

		// process
		// - get all vehicles
		err = h.sv.UpdateSpeed(vehicle)
		if err != nil {
			if errors.Is(err, domain.ErrVehicleNotFound) {
				response.JSON(w, http.StatusNotFound, "Veículo não encontrado.")
				return
			}
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

		response.JSON(w, http.StatusOK, "Velocidade do veículo atualizada com sucesso")
	}
}

func (h *VehicleDefault) GetByFuelType() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		t := chi.URLParam(r, "type")

		// process
		// - get all vehicles
		v, err := h.sv.GetByFuelType(t)
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

		if len(v) <= 0 {
			response.JSON(w, http.StatusNotFound, "Não foram encontrados veículos com esse tipo de combustível.")
			return
		}

		// response
		data := make(map[int]VehicleJSON)
		for key, value := range v {
			data[key] = VehicleJSON{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:        value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}
		response.JSON(w, http.StatusOK, data)
	}
}

func (h *VehicleDefault) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		idStr := chi.URLParam(r, "id")
		var id int
		fmt.Sscanf(idStr, "%d", &id)

		// process
		// - get all vehicles
		err := h.sv.Delete(id)
		if err != nil {
			if errors.Is(err, domain.ErrVehicleNotFound) {
				response.JSON(w, http.StatusNotFound, "Veículo não encontrado.")
			}
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

		response.JSON(w, http.StatusNoContent, "Veículo removido com sucesso.")
	}
}

func (h *VehicleDefault) GetByTransmissionType() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		t := chi.URLParam(r, "type")

		// process
		// - get all vehicles
		v, err := h.sv.GetByTransmissionType(t)
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

		if len(v) <= 0 {
			response.JSON(w, http.StatusNotFound, "Não foram encontrados veículos com esse tipo de transmissão.")
			return
		}

		// response
		data := make(map[int]VehicleJSON)
		for key, value := range v {
			data[key] = VehicleJSON{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:        value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}
		response.JSON(w, http.StatusOK, data)
	}
}

func (h *VehicleDefault) UpdateFuel() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		idStr := chi.URLParam(r, "id")
		var id int
		fmt.Sscanf(idStr, "%d", id)

		var reqBody VehicleJSON
		err := json.NewDecoder(r.Body).Decode(&reqBody)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, "Tipo de combustível malformado ou não suportado.")
		}

		vehicle := internal.Vehicle{
			Id: reqBody.ID,
			VehicleAttributes: internal.VehicleAttributes{
				Brand:           reqBody.Brand,
				Model:           reqBody.Model,
				Registration:    reqBody.Registration,
				Color:           reqBody.Color,
				FabricationYear: reqBody.FabricationYear,
				Capacity:        reqBody.Capacity,
				MaxSpeed:        reqBody.MaxSpeed,
				FuelType:        reqBody.FuelType,
				Transmission:    reqBody.Transmission,
				Weight:          reqBody.Weight,
				Dimensions: internal.Dimensions{
					Height: reqBody.Height,
					Length: reqBody.Length,
					Width:  reqBody.Width,
				},
			},
		}

		// process
		// - get all vehicles
		err = h.sv.UpdateFuel(vehicle)
		if err != nil {
			if errors.Is(err, domain.ErrVehicleNotFound) {
				response.JSON(w, http.StatusNotFound, "Veículo não encontrado.")
				return
			}
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

		response.JSON(w, http.StatusOK, "Tipo do combustível atualizado com sucesso.")
	}
}

func (h *VehicleDefault) GetAverageCapacityByBrand() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		brand := chi.URLParam(r, "brand")

		// process
		// - get all vehicles
		ac, err := h.sv.GetAverageCapacityByBrand(brand)
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

		if ac <= 0 {
			response.JSON(w, http.StatusNotFound, nil)
			return
		}

		response.JSON(w, http.StatusOK, ac)
	}
}

func (h *VehicleDefault) GetByDimensions() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		lengthRange := r.URL.Query().Get("length")
		widthRange := r.URL.Query().Get("width")

		minLength, maxLength, err1 := parseRange(lengthRange)
		minWidth, maxWidth, err2 := parseRange(widthRange)

		if err1 != nil || err2 != nil {
			response.JSON(w, http.StatusBadRequest, "Dimensões não puderam ser lidas.")
		}

		// process
		// - get all vehicles
		v, err := h.sv.GetByDimentions(minLength, maxLength, minWidth, maxWidth)
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

		if len(v) <= 0 {
			response.JSON(w, http.StatusNotFound, "Não foram encontrados veículos com essas dimensões.")
		}

		// response
		data := make(map[int]VehicleJSON)
		for key, value := range v {
			data[key] = VehicleJSON{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:        value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}
		response.JSON(w, http.StatusOK, data)
	}
}

func parseRange(rangeStr string) (min float64, max float64, err error) {
	parts := strings.Split(rangeStr, "-")
	if len(parts) != 2 {
		return 0, 0, domain.ErrInvalidRangeFormats
	}

	min, err = strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return 0, 0, err
	}

	max, err = strconv.ParseFloat(parts[1], 64)
	if err != nil {
		return 0, 0, err
	}

	return
}

func (h *VehicleDefault) GetVehiclesByWeightRange() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		minStr := r.URL.Query().Get("min")
		maxStr := r.URL.Query().Get("max")

		min, err1 := strconv.ParseFloat(minStr, 64)
		max, err2 := strconv.ParseFloat(maxStr, 64)
		if err1 != nil || err2 != nil {
			response.JSON(w, http.StatusBadRequest, "Parâmetros enviados não estão corretos")
		}

		// process
		// - get all vehicles
		v, err := h.sv.GetVehiclesByWeightRange(min, max)
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

		if len(v) <= 0 {
			response.JSON(w, http.StatusNotFound, "Não foram encontrados veículos nessa faixa de peso.")
			return
		}

		// response
		data := make(map[int]VehicleJSON)
		for key, value := range v {
			data[key] = VehicleJSON{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:        value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}
