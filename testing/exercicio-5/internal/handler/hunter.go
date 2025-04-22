package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"testdoubles/internal/hunter"
	"testdoubles/internal/positioner"
	"testdoubles/internal/prey"
	"testdoubles/platform/web/response"
)

// NewHunter returns a new Hunter handler.
func NewHunter(ht hunter.Hunter, pr prey.Prey) *Hunter {
	return &Hunter{ht: ht, pr: pr}
}

// Hunter returns handlers to manage hunting.
type Hunter struct {
	// ht is the Hunter interface that this handler will use
	ht hunter.Hunter
	// pr is the Prey interface that the hunter will hunt
	pr prey.Prey
}

// RequestBodyConfigPrey is an struct to configure the prey for the hunter in JSON format.
type RequestBodyConfigPrey struct {
	Speed    float64              `json:"speed"`
	Position *positioner.Position `json:"position"`
}

// ConfigurePrey configures the prey for the hunter.
func (h *Hunter) ConfigurePrey() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		var reqBody RequestBodyConfigPrey
		json.NewDecoder(r.Body).Decode(&reqBody)

		// process
		h.pr.Configure(reqBody.Speed, reqBody.Position)

		// response
		res := map[string]interface{}{"position": h.pr.GetPosition(), "speed": h.pr.GetSpeed()}
		response.JSON(w, http.StatusOK, res)
	}
}

// RequestBodyConfigHunter is an struct to configure the hunter in JSON format.
type RequestBodyConfigHunter struct {
	Speed    float64              `json:"speed"`
	Position *positioner.Position `json:"position"`
}

// ConfigureHunter configures the hunter.
func (h *Hunter) ConfigureHunter() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		var reqBody RequestBodyConfigHunter
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			response.Error(w, http.StatusBadRequest, "JSON inválido!")
			return
		}

		// process
		h.ht.Configure(reqBody.Speed, reqBody.Position)

		// response
		res := map[string]interface{}{"position": h.pr.GetPosition(), "speed": h.pr.GetSpeed()}
		response.JSON(w, http.StatusOK, res)
	}
}

// Hunt hunts the prey.
func (h *Hunter) Hunt() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request

		// process
		duration, err := h.ht.Hunt(h.pr)
		if err != nil {
			if !errors.Is(err, hunter.ErrCanNotHunt) {
				response.Error(w, http.StatusInternalServerError, "internal server error")
				return
			}
		}

		// response
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "caça concluída",
			"data": map[string]any{
				"success":  err == nil,
				"duration": duration,
			},
		})
	}
}
