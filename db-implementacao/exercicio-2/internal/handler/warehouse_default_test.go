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

	"github.com/DATA-DOG/go-txdb"
	"github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/require"
)

func init() {
	cfg := mysql.Config{
		User:   "root",
		Passwd: "root",
		Net:    "tcp",
		Addr:   "localhost:3306",
		DBName: "my_db",
	}

	txdb.Register("txdb", "mysql", cfg.FormatDSN())
}

type WarehousesResponse struct {
	Data    []internal.Warehouse `json:"data"`
	Message string               `json:"message"`
}

func TestHandlerWarehouseGetAll(t *testing.T) {
	t.Run("success - get all warehouses", func(t *testing.T) {
		// arrange
		db, err := sql.Open("txdb", "my_db")
		require.NoError(t, err)
		defer db.Close()

		rp := repository.NewWarehousesMySQL(db)
		hd := handler.NewWarehousesDefault(rp)
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

func TestHandlerWarehouseGetOne(t *testing.T) {
	t.Run("success - get a warehouses", func(t *testing.T) {
		// arrange
		db, err := sql.Open("txdb", "my_db")
		require.NoError(t, err)
		defer db.Close()

		rp := repository.NewWarehousesMySQL(db)
		hd := handler.NewWarehousesDefault(rp)

		// act
		request := httptest.NewRequest("GET", "/warehouses/1", nil)
		request = testutils.WithUrlParam(t, request, "id", "1")
		response := httptest.NewRecorder()
		hd.GetOne()(response, request)

		// assert
		expectedCode := http.StatusOK
		expectedBody := `{"data":{"address":"221 Baker Street", "capacity":100, "id":1, "name":"Main Warehouse", "telephone":"4555666"}, "message":"warehouse found"}`
		expectedHeader := http.Header{
			"Content-Type": []string{"application/json"},
		}

		require.Equal(t, expectedCode, response.Code)
		require.JSONEq(t, expectedBody, response.Body.String())
		require.Equal(t, expectedHeader, response.Header())
	})
}

func TestHandlerWarehouseCreate(t *testing.T) {
	t.Run("success - create a warehouse", func(t *testing.T) {
		// arrange
		db, err := sql.Open("txdb", "my_db")
		require.NoError(t, err)
		defer db.Close()

		rp := repository.NewWarehousesMySQL(db)
		hd := handler.NewWarehousesDefault(rp)
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
		expectedBody := `{"data":{"address":"", "capacity":0, "id":8, "name":"Teste", "telephone":""}, "message":"warehouse created"}`
		expectedHeader := http.Header{
			"Content-Type": []string{"application/json"},
		}

		require.Equal(t, expectedCode, response.Code)
		require.JSONEq(t, expectedBody, response.Body.String())
		require.Equal(t, expectedHeader, response.Header())
	})
}
