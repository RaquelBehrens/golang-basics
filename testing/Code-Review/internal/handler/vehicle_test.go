package handler_test

import (
	"app/internal"
	"app/internal/handler"
	"app/internal/repository"
	"app/internal/service"
	"app/testutils"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindByColorAndYear(t *testing.T) {
	t.Run("find vehicle by color and year - success", func(t *testing.T) {
		// given
		rp := repository.NewRepositoryMock()
		vehicles := map[int]internal.Vehicle{
			1: {
				Id: 1,
				VehicleAttributes: internal.VehicleAttributes{
					Brand:           "Ford",
					Model:           "Ka",
					Registration:    "ABC1234",
					Color:           "Prata",
					FabricationYear: 2015,
					Capacity:        5,
					MaxSpeed:        190,
					FuelType:        "Gasolina",
					Transmission:    "Manual",
					Weight:          1050,
					Dimensions: internal.Dimensions{
						Height: 1.52,
						Length: 3.89,
						Width:  1.70,
					},
				},
			},
			2: {
				Id: 2,
				VehicleAttributes: internal.VehicleAttributes{
					Brand:           "Chevrolet",
					Model:           "Onix",
					Registration:    "XYZ5678",
					Color:           "Preto",
					FabricationYear: 2018,
					Capacity:        5,
					MaxSpeed:        180,
					FuelType:        "Flex",
					Transmission:    "Automático",
					Weight:          1070,
					Dimensions: internal.Dimensions{
						Height: 1.48,
						Length: 3.93,
						Width:  1.70,
					},
				},
			},
		}
		rp.Mock.On("FindByColorAndYear", "Prata", 2015).Return(vehicles, nil)
		sr := service.NewServiceVehicleDefault(rp)
		hd := handler.NewHandlerVehicle(sr)

		expectedBodyMap := map[string]interface{}{
			"data":    vehicles,
			"message": "vehicles found",
		}
		params := map[string]string{
			"color": "Prata",
			"year":  "2015",
		}

		// when
		req := httptest.NewRequest("GET", "/color/{color}/year/{year}", nil)
		// req = testutils.WithUrlParam(t, req, "color", "Prata")
		// req = testutils.WithUrlParam(t, req, "year", "2015")
		req = testutils.WithUrlParamst(t, req, params)
		res := httptest.NewRecorder()
		hd.FindByColorAndYear()(res, req)

		// then
		expectedCode := http.StatusOK
		expectedHeader := http.Header{"Content-Type": []string{"application/json; charset=utf-8"}}
		expectedBody, _ := json.Marshal(expectedBodyMap)

		require.Equal(t, expectedCode, res.Code)
		require.Equal(t, expectedHeader, res.Header())
		require.JSONEq(t, string(expectedBody), res.Body.String())
		// require.NotEmpty(t, res.Body.String(), "Response body should not be empty")

		// var responseBodyMap map[string]interface{}
		// err := json.Unmarshal(res.Body.Bytes(), &responseBodyMap)
		// require.NoError(t, err)

		// for key, expectedValue := range expectedBodyMap {
		// 	require.Equal(t, expectedValue, responseBodyMap[key])
		// }
	})

	t.Run("find vehicle by year - not found", func(t *testing.T) {
		// given
		rp := repository.NewRepositoryMock()
		vehicles := map[int]internal.Vehicle{}
		rp.Mock.On("FindByColorAndYear", "Prata", 2015).Return(vehicles, nil)
		sr := service.NewServiceVehicleDefault(rp)
		hd := handler.NewHandlerVehicle(sr)

		expectedBodyMap := map[string]string{
			"message": "no vehicles found",
			"status":  "Not Found",
		}

		// when
		req := httptest.NewRequest("GET", "/color/{color}/year/{year}", nil)
		params := map[string]string{
			"color": "Prata",
			"year":  "2015",
		}
		req = testutils.WithUrlParamst(t, req, params)
		res := httptest.NewRecorder()
		hd.FindByColorAndYear()(res, req)

		// then
		expectedCode := http.StatusNotFound
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedBody, _ := json.Marshal(expectedBodyMap)

		require.Equal(t, expectedCode, res.Code)
		require.Equal(t, expectedHeader, res.Header())
		require.JSONEq(t, string(expectedBody), res.Body.String())

	})
}

func TestFindByBrandAndYearRange(t *testing.T) {
	t.Run("find by brand and year range - ok", func(t *testing.T) {
		rp := repository.NewRepositoryMock()
		vehicles := map[int]internal.Vehicle{
			1: {
				Id: 1,
				VehicleAttributes: internal.VehicleAttributes{
					Brand:           "Teste",
					Model:           "Teste",
					Registration:    "Teste",
					Color:           "Teste",
					FabricationYear: 2015,
					Capacity:        5,
					MaxSpeed:        200.0,
					FuelType:        "Teste",
					Transmission:    "Teste",
					Weight:          100.0,
					Dimensions: internal.Dimensions{
						Height: 1.0,
						Length: 1.0,
						Width:  1.0},
				},
			},
			2: {
				Id: 2,
				VehicleAttributes: internal.VehicleAttributes{
					Brand:           "Teste",
					Model:           "Teste",
					Registration:    "Teste",
					Color:           "Teste",
					FabricationYear: 2016,
					Capacity:        5,
					MaxSpeed:        200.0,
					FuelType:        "Teste",
					Transmission:    "Teste",
					Weight:          100.0,
					Dimensions: internal.Dimensions{
						Height: 1.0,
						Length: 1.0,
						Width:  1.0},
				},
			},
		}
		rp.Mock.On("FindByBrandAndYearRange", "Teste", 2000, 2020).Return(vehicles, nil)
		sr := service.NewServiceVehicleDefault(rp)
		hd := handler.NewHandlerVehicle(sr)

		expectedBodyMap := map[string]interface{}{
			"message": "vehicles found",
			"data":    vehicles,
		}

		// when
		req := httptest.NewRequest("GET", "/brand/{brand}/between/{start_year}/{end_year}", nil)
		params := map[string]string{"brand": "Teste", "start_year": "2000", "end_year": "2020"}
		req = testutils.WithUrlParamst(t, req, params)
		res := httptest.NewRecorder()
		hd.FindByBrandAndYearRange()(res, req)

		// then
		expectedCode := http.StatusOK
		expectedHeader := http.Header{"Content-Type": []string{"application/json; charset=utf-8"}}
		expectedBody, _ := json.Marshal(expectedBodyMap)

		require.Equal(t, expectedCode, res.Code)
		require.Equal(t, expectedHeader, res.Header())
		require.JSONEq(t, string(expectedBody), res.Body.String())
	})

	t.Run("find by brand and year range - not found", func(t *testing.T) {
		// given
		rp := repository.NewRepositoryMock()
		vehicles := map[int]internal.Vehicle{}
		rp.Mock.On("FindByBrandAndYearRange", "Teste", 2000, 2020).Return(vehicles, nil)

		sr := service.NewServiceVehicleDefault(rp)
		hd := handler.NewHandlerVehicle(sr)

		expectedBodyMap := map[string]string{
			"message": "no vehicles found",
			"status":  "Not Found",
		}

		// when
		req := httptest.NewRequest("GET", "/brand/{brand}/between/{start_year}/{end_year}", nil)
		params := map[string]string{"brand": "Teste", "start_year": "2000", "end_year": "2020"}
		req = testutils.WithUrlParamst(t, req, params)
		res := httptest.NewRecorder()
		hd.FindByBrandAndYearRange()(res, req)

		// then
		expectedCode := http.StatusNotFound
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedBody, _ := json.Marshal(expectedBodyMap)

		require.Equal(t, expectedCode, res.Code)
		require.Equal(t, expectedHeader, res.Header())
		require.JSONEq(t, string(expectedBody), res.Body.String())
	})
}

func TestAverageMaxSpeedByBrand(t *testing.T) {
	t.Run("average max speed by brand - ok", func(t *testing.T) {
		// given
		rp := repository.NewRepositoryMock()
		vehicles := map[int]internal.Vehicle{
			1: {
				Id: 1,
				VehicleAttributes: internal.VehicleAttributes{
					Brand:           "Teste",
					Model:           "Teste",
					Registration:    "Teste",
					Color:           "Teste",
					FabricationYear: 1,
					Capacity:        1,
					MaxSpeed:        100.0,
					FuelType:        "Teste",
					Transmission:    "Teste",
					Weight:          1.0,
					Dimensions: internal.Dimensions{
						Height: 1.0,
						Length: 1.0,
						Width:  1.0,
					},
				},
			},
		}
		rp.Mock.On("FindByBrand", "Teste").Return(vehicles, nil)
		sr := service.NewServiceVehicleDefault(rp)
		hd := handler.NewHandlerVehicle(sr)

		expectedBodyMap := map[string]interface{}{
			"data":    100,
			"message": "average max speed found",
		}

		// when
		req := httptest.NewRequest("GET", "/average_speed/brand/{brand}", nil)
		req = testutils.WithUrlParam(t, req, "brand", "Teste")
		res := httptest.NewRecorder()
		hd.AverageMaxSpeedByBrand()(res, req)

		// then
		expectedCode := http.StatusOK
		expectedHeader := http.Header{"Content-Type": []string{"application/json; charset=utf-8"}}
		expectedBody, _ := json.Marshal(expectedBodyMap)

		require.Equal(t, expectedCode, res.Code)
		require.Equal(t, expectedHeader, res.Header())
		require.JSONEq(t, string(expectedBody), res.Body.String())
	})

	t.Run("average max speed by brand - not found", func(t *testing.T) {
		// given
		rp := repository.NewRepositoryMock()
		vehicles := map[int]internal.Vehicle{}
		rp.Mock.On("FindByBrand", "Teste").Return(vehicles, nil)
		sr := service.NewServiceVehicleDefault(rp)
		hd := handler.NewHandlerVehicle(sr)

		expectedBodyMap := map[string]interface{}{"message": "vehicles not found", "status": "Not Found"}

		// when
		req := httptest.NewRequest("GET", "/average_speed/brand/{brand}", nil)
		req = testutils.WithUrlParam(t, req, "brand", "Teste")
		res := httptest.NewRecorder()
		hd.AverageMaxSpeedByBrand()(res, req)

		// then
		expectedCode := http.StatusNotFound
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedBody, _ := json.Marshal(expectedBodyMap)

		require.Equal(t, expectedCode, res.Code)
		require.Equal(t, expectedHeader, res.Header())
		require.JSONEq(t, string(expectedBody), res.Body.String())
	})
}

func TestAverageCapacityByBrand(t *testing.T) {
	t.Run("average capacity by brand - success", func(t *testing.T) {
		// given
		rp := repository.NewRepositoryMock()
		vehicles := map[int]internal.Vehicle{
			1: {
				Id: 1,
				VehicleAttributes: internal.VehicleAttributes{
					Brand:           "Teste",
					Model:           "Teste",
					Registration:    "Teste",
					Color:           "Teste",
					FabricationYear: 1,
					Capacity:        5,
					MaxSpeed:        100.0,
					FuelType:        "Teste",
					Transmission:    "Teste",
					Weight:          100.0,
					Dimensions: internal.Dimensions{
						Height: 100.0,
						Length: 100.0,
						Width:  100.0,
					},
				},
			},
			2: {
				Id: 2,
				VehicleAttributes: internal.VehicleAttributes{
					Brand:           "Teste",
					Model:           "Teste",
					Registration:    "Teste",
					Color:           "Teste",
					FabricationYear: 1,
					Capacity:        4,
					MaxSpeed:        100.0,
					FuelType:        "Teste",
					Transmission:    "Teste",
					Weight:          100.0,
					Dimensions: internal.Dimensions{
						Height: 100.0,
						Length: 100.0,
						Width:  100.0,
					},
				},
			},
		}
		rp.Mock.On("FindByBrand", "Teste").Return(vehicles, nil)
		sr := service.NewServiceVehicleDefault(rp)
		hd := handler.NewHandlerVehicle(sr)

		expectedBodyMap := map[string]interface{}{"data": 4, "message": "average capacity found"}

		// when
		req := httptest.NewRequest("GET", "/average_capacity/brand/{brand}", nil)
		req = testutils.WithUrlParam(t, req, "brand", "Teste")
		res := httptest.NewRecorder()
		hd.AverageCapacityByBrand()(res, req)

		// then
		expectedCode := http.StatusOK
		expectedHeader := http.Header{"Content-Type": []string{"application/json; charset=utf-8"}}
		expectedBody, _ := json.Marshal(expectedBodyMap)

		require.Equal(t, expectedCode, res.Code)
		require.Equal(t, expectedHeader, res.Header())
		require.JSONEq(t, string(expectedBody), res.Body.String())
	})

	t.Run("average capacity by brand - not found", func(t *testing.T) {
		// given
		rp := repository.NewRepositoryMock()
		vehicles := map[int]internal.Vehicle{}
		rp.Mock.On("FindByBrand", "Teste").Return(vehicles, nil)
		sr := service.NewServiceVehicleDefault(rp)
		hd := handler.NewHandlerVehicle(sr)

		expectedBodyMap := map[string]interface{}{"message": "vehicles not found", "status": "Not Found"}

		// when
		req := httptest.NewRequest("GET", "/average_capacity/brand/{brand}", nil)
		req = testutils.WithUrlParam(t, req, "brand", "Teste")
		res := httptest.NewRecorder()
		hd.AverageCapacityByBrand()(res, req)

		// then
		expectedCode := http.StatusNotFound
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedBody, _ := json.Marshal(expectedBodyMap)

		require.Equal(t, expectedCode, res.Code)
		require.Equal(t, expectedHeader, res.Header())
		require.JSONEq(t, string(expectedBody), res.Body.String())
	})
}

func TestSearchByWeightRange(t *testing.T) {
	t.Run("search by weight range - ok", func(t *testing.T) {
		// given
		rp := repository.NewRepositoryMock()
		vehicles := map[int]internal.Vehicle{
			1: {
				Id: 1,
				VehicleAttributes: internal.VehicleAttributes{
					Brand:           "Teste",
					Model:           "Teste",
					Registration:    "Teste",
					Color:           "Teste",
					FabricationYear: 2000,
					Capacity:        5,
					MaxSpeed:        100.0,
					FuelType:        "Teste",
					Transmission:    "Teste",
					Weight:          230.0,
					Dimensions: internal.Dimensions{
						Height: 100.0,
						Length: 100.0,
						Width:  100.0,
					},
				},
			},
		}
		rp.Mock.On("FindByWeightRange", 200.0, 300.0).Return(vehicles, nil)
		sr := service.NewServiceVehicleDefault(rp)
		hd := handler.NewHandlerVehicle(sr)

		// when
		req := httptest.NewRequest("GET", "/weight", nil)
		params := map[string]string{
			"weight_min": "200",
			"weight_max": "300",
		}
		req = testutils.WithQueryParams(t, req, params)
		res := httptest.NewRecorder()
		hd.SearchByWeightRange()(res, req)

		// then
		expectedBodyMap := map[string]interface{}{"data": vehicles, "message": "vehicles found"}

		expectedCode := http.StatusOK
		expectedHeader := http.Header{"Content-Type": []string{"application/json; charset=utf-8"}}
		expectedBody, _ := json.Marshal(expectedBodyMap)

		require.Equal(t, expectedCode, res.Code)
		require.Equal(t, expectedHeader, res.Header())
		require.JSONEq(t, string(expectedBody), res.Body.String())
	})

	t.Run("search by weight range - not found", func(t *testing.T) {
		// given
		rp := repository.NewRepositoryMock()
		vehicles := map[int]internal.Vehicle{}
		rp.Mock.On("FindByWeightRange", 200.0, 300.0).Return(vehicles, nil)
		sr := service.NewServiceVehicleDefault(rp)
		hd := handler.NewHandlerVehicle(sr)

		// when
		req := httptest.NewRequest("GET", "/weight", nil)
		params := map[string]string{
			"weight_min": "200",
			"weight_max": "300",
		}
		req = testutils.WithQueryParams(t, req, params)
		res := httptest.NewRecorder()
		hd.SearchByWeightRange()(res, req)

		// then
		expectedBodyMap := map[string]interface{}{"message": "no vehicles found", "status": "Not Found"}

		expectedCode := http.StatusNotFound
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedBody, _ := json.Marshal(expectedBodyMap)

		require.Equal(t, expectedCode, res.Code)
		require.Equal(t, expectedHeader, res.Header())
		require.JSONEq(t, string(expectedBody), res.Body.String())
	})
}
