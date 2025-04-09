package handler_test

import (
	"app/internal"
	"app/internal/handler"
	"app/internal/repository"
	"app/internal/testutils"
	"database/sql"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

type ProductsResponse struct {
	Data    []internal.Product `json:"data"`
	Message string             `json:"message"`
}

func TestHandlerProductGetAll(t *testing.T) {
	t.Run("success - get all products", func(t *testing.T) {
		// arrange
		db, err := sql.Open("txdb", "my_db")
		require.NoError(t, err)
		defer db.Close()

		rp := repository.NewProductsMySQL(db)
		hd := handler.NewProductsDefault(rp)
		hdFunc := hd.GetAll()

		// act
		request := &http.Request{}
		response := httptest.NewRecorder()
		hdFunc(response, request)

		// assert
		expectedCode := http.StatusOK
		// expectedBody := `{"message":"ID inválido!", "status":"Bad Request"}`
		expectedHeader := http.Header{
			"Content-Type": []string{"application/json"},
		}

		require.Equal(t, expectedCode, response.Code)
		// require.JSONEq(t, expectedBody, response.Body.String())
		require.Equal(t, expectedHeader, response.Header())
	})
}

func TestHandlerProductGetOne(t *testing.T) {
	t.Run("success - get a products", func(t *testing.T) {
		// arrange
		db, err := sql.Open("txdb", "my_db")
		require.NoError(t, err)
		defer db.Close()

		rp := repository.NewProductsMySQL(db)
		hd := handler.NewProductsDefault(rp)

		// act
		request := httptest.NewRequest("GET", "/products/1", nil)
		request = testutils.WithUrlParam(t, request, "id", "1")
		response := httptest.NewRecorder()
		hd.GetOne()(response, request)

		// assert
		expectedCode := http.StatusOK
		expectedBody := `{"data":{"code_value":"0009-1111", "expiration":"2022-01-08", "id":1, "is_published":false, "name":"Corn Shoots", "price":23.27, "quantity":244}, "message":"product found"}`
		expectedHeader := http.Header{
			"Content-Type": []string{"application/json"},
		}

		require.Equal(t, expectedCode, response.Code)
		require.JSONEq(t, expectedBody, response.Body.String())
		require.Equal(t, expectedHeader, response.Header())
	})
}

func TestHandlerProductCreate(t *testing.T) {
	t.Run("success - create a product", func(t *testing.T) {
		// arrange
		db, err := sql.Open("txdb", "my_db")
		require.NoError(t, err)
		defer db.Close()

		rp := repository.NewProductsMySQL(db)
		hd := handler.NewProductsDefault(rp)
		hdFunc := hd.Create()

		// act
		request := &http.Request{
			Body: io.NopCloser(strings.NewReader(
				`{"id": 1, "name":"Teste", "quantity": 1, "code_value": "", "is_published": false, "expiration": "2025-01-01", "price": 1.1, "id_warehouse": 1}`,
			)),
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
		}
		response := httptest.NewRecorder()
		hdFunc(response, request)

		// assert
		expectedCode := http.StatusCreated
		expectedBody := `{"data":{"id":415,"name":"Teste","quantity":1,"code_value":"","is_published":false,"expiration":"2025-01-01","price":1.1},"message":"product created"}`
		expectedHeader := http.Header{
			"Content-Type": []string{"application/json"},
		}

		require.Equal(t, expectedCode, response.Code)
		require.JSONEq(t, expectedBody, response.Body.String())
		require.Equal(t, expectedHeader, response.Header())
	})
}
