package handler_test

import (
	"app/internal/handler"
	"app/internal/repository"
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUpdateAllTotals(t *testing.T) {
	t.Run("success getting invoices by condition", func(t *testing.T) {
		db, err := sql.Open("txdb", "")
		require.NoError(t, err)
		defer db.Close()

		rp := repository.NewInvoicesMySQL(db)
		hd := handler.NewInvoicesDefault(rp)

		request := httptest.NewRequest("GET", "/invoices/updateAllTotals", nil)
		response := httptest.NewRecorder()
		hd.UpdateAllTotals()(response, request)

		expectedCode := http.StatusOK
		expectedHeader := http.Header{
			"Content-Type": []string{"application/json"},
		}

		require.Equal(t, expectedCode, response.Code)
		require.Equal(t, expectedHeader, response.Header())

	})
}
