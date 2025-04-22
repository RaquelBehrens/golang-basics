package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testdoubles/internal/handler"
	"testdoubles/internal/hunter"
	"testdoubles/internal/positioner"
	"testdoubles/internal/prey"
	"testdoubles/internal/simulator"
	"testing"

	"github.com/stretchr/testify/require"
)

func setupHandler() *handler.Hunter {
	ps := positioner.NewPositionerDefault()
	sm := simulator.NewCatchSimulatorDefault(&simulator.ConfigCatchSimulatorDefault{
		Positioner: ps,
	})
	ht := hunter.NewWhiteShark(hunter.ConfigWhiteShark{
		Speed:     0.0,
		Position:  &positioner.Position{X: 0.0, Y: 0.0, Z: 0.0},
		Simulator: sm,
	})
	pr := prey.NewTuna(0.0, &positioner.Position{X: 0.0, Y: 0.0, Z: 0.0})
	return handler.NewHunter(ht, pr)
}

func TestHunter_ConfigurePrey(t *testing.T) {
	t.Run("success configuring prey", func(t *testing.T) {
		// given
		hd := setupHandler()
		newPrey := handler.RequestBodyConfigPrey{Speed: 0, Position: &positioner.Position{X: 0, Y: 0, Z: 0}}
		jsonNewPrey, _ := json.Marshal(newPrey)

		// when
		req := httptest.NewRequest("POST", "/configure-prey", bytes.NewReader(jsonNewPrey))
		res := httptest.NewRecorder()
		hd.ConfigurePrey()(res, req)

		// then
		expectedCode := http.StatusOK
		expectedBody := `{"speed":0,"position":{"X":0,"Y":0,"Z":0}}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, expectedHeader, res.Header())
	})

}

func TestHunter_ConfigureHunter(t *testing.T) {
	t.Run("invalid json while configuring hunter", func(t *testing.T) {
		// given
		hd := setupHandler()
		newHunter := handler.RequestBodyConfigHunter{Speed: 0, Position: &positioner.Position{X: 0, Y: 0, Z: 0}}
		jsonNewHunter, _ := json.Marshal(newHunter)

		// when
		req := httptest.NewRequest("POST", "/configure-hunter", bytes.NewReader(jsonNewHunter))
		res := httptest.NewRecorder()
		hd.ConfigureHunter()(res, req)

		// then
		expectedCode := http.StatusBadRequest
		expectedBody := `"Json inválido!"`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, expectedHeader, res.Header())
	})

}
