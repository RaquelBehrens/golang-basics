package handler_test

import (
	"app/internal"
	"app/internal/handler"
	"app/internal/repository"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProducts_Get(t *testing.T) {
	t.Run("case 1: successfuly get a product", func(t *testing.T) {
		// given
		ht := repository.NewProductsMock()
		product := map[int]internal.Product{
			1: {Id: 1, ProductAttributes: internal.ProductAttributes{Description: "Test1", Price: 1.0, SellerId: 1}},
		}
		ht.FuncSearchProducts = func(query internal.ProductQuery) (p map[int]internal.Product, err error) {
			return product, nil
		}
		hd := handler.NewProductsDefault(ht)

		// when
		req := httptest.NewRequest("GET", "/", nil)
		q := req.URL.Query()
		q.Add("id", "1")
		req.URL.RawQuery = q.Encode()
		res := httptest.NewRecorder()
		hd.Get()(res, req)

		// then
		expectedCode := http.StatusOK
		expectedBody := `{
			"data": {"1":{"description":"Test1", "id":1, "price":1, "seller_id":1}}, 
			"message":"success"
		}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedCallHunt := 1

		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, expectedHeader, res.Header())
		require.Equal(t, expectedCallHunt, ht.Spy.SearchProducts)
	})

	t.Run("case 2: successfuly get a product - empty", func(t *testing.T) {
		// given
		ht := repository.NewProductsMock()
		product := map[int]internal.Product{}
		ht.FuncSearchProducts = func(query internal.ProductQuery) (p map[int]internal.Product, err error) {
			return product, nil
		}
		hd := handler.NewProductsDefault(ht)

		// when
		req := httptest.NewRequest("GET", "/", nil)
		q := req.URL.Query()
		q.Add("id", "1")
		req.URL.RawQuery = q.Encode()
		res := httptest.NewRecorder()
		hd.Get()(res, req)

		// then
		expectedCode := http.StatusOK
		expectedBody := `{
			"data": {}, 
			"message":"success"
		}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedCallHunt := 1

		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, expectedHeader, res.Header())
		require.Equal(t, expectedCallHunt, ht.Spy.SearchProducts)
	})

	t.Run("case 3: search all products - db empty", func(t *testing.T) {
		// given
		ht := repository.NewProductsMock()
		product := map[int]internal.Product{}
		ht.FuncSearchProducts = func(query internal.ProductQuery) (p map[int]internal.Product, err error) {
			return product, nil
		}
		hd := handler.NewProductsDefault(ht)

		// when
		req := httptest.NewRequest("GET", "/", nil)
		res := httptest.NewRecorder()
		hd.Get()(res, req)

		// then
		expectedCode := http.StatusOK
		expectedBody := `{
			"data": {}, 
			"message":"success"
		}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedCallHunt := 1

		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, expectedHeader, res.Header())
		require.Equal(t, expectedCallHunt, ht.Spy.SearchProducts)
	})

	t.Run("case 4: search all products - full", func(t *testing.T) {
		// given
		ht := repository.NewProductsMock()
		products := map[int]internal.Product{
			1: {Id: 1, ProductAttributes: internal.ProductAttributes{Description: "Test1", Price: 1.0, SellerId: 1}},
			2: {Id: 2, ProductAttributes: internal.ProductAttributes{Description: "Test2", Price: 2.0, SellerId: 2}},
			3: {Id: 3, ProductAttributes: internal.ProductAttributes{Description: "Test3", Price: 3.0, SellerId: 3}},
		}
		ht.FuncSearchProducts = func(query internal.ProductQuery) (p map[int]internal.Product, err error) {
			return products, nil
		}
		hd := handler.NewProductsDefault(ht)

		// when
		req := httptest.NewRequest("GET", "/", nil)
		res := httptest.NewRecorder()
		hd.Get()(res, req)

		// then
		expectedCode := http.StatusOK
		expectedBody := `{
			"data": {
				"1":{"description":"Test1", "id":1, "price":1, "seller_id":1},
				"2":{"description":"Test2", "id":2, "price":2, "seller_id":2},
				"3":{"description":"Test3", "id":3, "price":3, "seller_id":3}
				}, 
			"message":"success"
		}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedCallHunt := 1

		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, expectedHeader, res.Header())
		require.Equal(t, expectedCallHunt, ht.Spy.SearchProducts)
	})

	t.Run("case 5: error trying to get a product - invalid id", func(t *testing.T) {
		// given
		ht := repository.NewProductsMock()
		ht.FuncSearchProducts = func(query internal.ProductQuery) (p map[int]internal.Product, err error) {
			return nil, errors.New("internal error")
		}
		hd := handler.NewProductsDefault(ht)

		// when
		req := httptest.NewRequest("GET", "/", nil)
		q := req.URL.Query()
		q.Add("id", "1")
		req.URL.RawQuery = q.Encode()
		res := httptest.NewRecorder()
		hd.Get()(res, req)

		// then
		expectedCode := http.StatusInternalServerError
		expectedBody := `{
			"message":"internal error",
			"status":"Internal Server Error"
		}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedCallHunt := 1

		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, expectedHeader, res.Header())
		require.Equal(t, expectedCallHunt, ht.Spy.SearchProducts)
	})

	t.Run("case 6: error trying to get a product - internal server error", func(t *testing.T) {
		// given
		ht := repository.NewProductsMock()
		hd := handler.NewProductsDefault(ht)

		// when
		req := httptest.NewRequest("GET", "/", nil)
		q := req.URL.Query()
		q.Add("id", "lala")
		req.URL.RawQuery = q.Encode()
		res := httptest.NewRecorder()
		hd.Get()(res, req)

		// then
		expectedCode := http.StatusBadRequest
		expectedBody := `{
			"message":"invalid id",
			"status":"Bad Request"
		}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedCallHunt := 0

		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, expectedHeader, res.Header())
		require.Equal(t, expectedCallHunt, ht.Spy.SearchProducts)
	})
}
