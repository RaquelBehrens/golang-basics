package handler_test

import (
	"app/internal/handler"
	"app/internal/repository"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-txdb"
	"github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/require"
)

func init() {
	cfg := mysql.Config{
		User:   "sgbd",
		Passwd: "root",
		Net:    "tcp",
		Addr:   "localhost:3306",
		DBName: "fantasy_products",
	}

	txdb.Register("txdb", "mysql", cfg.FormatDSN())
}

func TestGetInvoicesByCondition(t *testing.T) {
	t.Run("success getting invoices by condition", func(t *testing.T) {
		db, err := sql.Open("txdb", "")
		require.NoError(t, err)
		defer db.Close()

		rp := repository.NewCustomersMySQL(db)
		hd := handler.NewCustomersDefault(rp)

		request := httptest.NewRequest("GET", "/invoicesByCondition", nil)
		response := httptest.NewRecorder()
		hd.GetInvoicesByCondition()(response, request)

		data := map[string]interface{}{
			"data": []handler.CustomerInvoiceByConditionJSON{
				{Condition: 0, Total: 605929.1},
				{Condition: 1, Total: 716792.33},
			},
			"message": "invoices found by condition",
		}

		expectedCode := http.StatusOK
		expectedHeader := http.Header{
			"Content-Type": []string{"application/json"},
		}
		expectedBody, err := json.Marshal(data)
		if err != nil {
			log.Fatalf("Error marshalling to JSON: %v", err)
		}

		require.Equal(t, expectedCode, response.Code)
		require.JSONEq(t, string(expectedBody), response.Body.String())
		require.Equal(t, expectedHeader, response.Header())

	})
}

func TestGetMostActiveCustomersByAmountSpent(t *testing.T) {
	t.Run("success getting most active customers by amount spent", func(t *testing.T) {
		db, err := sql.Open("txdb", "")
		require.NoError(t, err)
		defer db.Close()

		rp := repository.NewCustomersMySQL(db)
		hd := handler.NewCustomersDefault(rp)

		request := httptest.NewRequest("GET", "/customers/mostActiveCustomersByAmountSpent", nil)
		response := httptest.NewRecorder()
		hd.GetMostActiveCustomersByAmountSpent()(response, request)

		data := map[string]interface{}{
			"data": []handler.MostActiveCustomersByAmountSpentJSON{
				{Amount: 58513.55, FirstName: "Lannie", LastName: "Tortis"},
				{Amount: 48291.03, FirstName: "Jasen", LastName: "Crowcum"},
				{Amount: 43590.75, FirstName: "Elvina", LastName: "Ovell"},
				{Amount: 40792.06, FirstName: "Lazaro", LastName: "Anstis"},
				{Amount: 39786.79, FirstName: "Wilden", LastName: "Oaten"}},
			"message": "most active customers found by condition",
		}

		expectedCode := http.StatusOK
		expectedHeader := http.Header{
			"Content-Type": []string{"application/json"},
		}
		expectedBody, err := json.Marshal(data)
		if err != nil {
			log.Fatalf("Error marshalling to JSON: %v", err)
		}

		require.Equal(t, expectedCode, response.Code)
		require.JSONEq(t, string(expectedBody), response.Body.String())
		require.Equal(t, expectedHeader, response.Header())

	})
}
