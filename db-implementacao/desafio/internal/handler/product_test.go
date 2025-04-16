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

	"github.com/stretchr/testify/require"
)

func TestGetBestSellingProducts(t *testing.T) {
	t.Run("success getting invoices by condition", func(t *testing.T) {
		db, err := sql.Open("txdb", "")
		require.NoError(t, err)
		defer db.Close()

		rp := repository.NewProductsMySQL(db)
		hd := handler.NewProductsDefault(rp)

		request := httptest.NewRequest("GET", "/invoicesByCondition", nil)
		response := httptest.NewRecorder()
		hd.GetBestSellingProducts()(response, request)

		data := map[string]interface{}{
			"data": []handler.BestSellingProductsJSON{
				{Description: "Vinegar - Raspberry", Total: 660},
				{Description: "Flour - Corn, Fine", Total: 521},
				{Description: "Cookie - Oatmeal", Total: 467},
				{Description: "Pepper - Red Chili", Total: 439},
				{Description: "Chocolate - Milk Coating", Total: 436},
			},
			"message": "best selling products found",
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
