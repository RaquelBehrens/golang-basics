package handlers_test

import (
	"exercicio/cmd/http/handlers"
	"exercicio/internal/domain"
	"exercicio/internal/products"
	"exercicio/internal/products/testutils"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
)

func setupHandler() *handlers.ProductHandler {
	godotenv.Load("../../../.env.example")

	db := map[int]domain.Product{
		1: {ID: 1, Name: "Teste 1", Quantity: 1, CodeValue: "T1", IsPublished: true, Expiration: time.Time{}, Price: 1.00},
		2: {ID: 2, Name: "Teste 2", Quantity: 2, CodeValue: "T2", IsPublished: false, Expiration: time.Time{}, Price: 2.00},
	}
	storage := products.NewStorage("filePath")
	repo := products.NewProductRepository(storage, db)
	srv := products.NewProductService(repo)
	return handlers.NewProductHandler(srv)
}

func TestProductsGetAll(t *testing.T) {
	t.Run("success to get products", func(t *testing.T) {
		// given
		hd := setupHandler()

		// when
		req := httptest.NewRequest("GET", "/products", nil)
		res := httptest.NewRecorder()
		hd.GetAll()(res, req)

		// then
		expectedCode := http.StatusOK
		expectedBody := `{"1":{"id":1,"name":"Teste 1","quantity":1,"codeValue":"T1","isPublished":true,"expiration":"0001-01-01T00:00:00Z","price":1},
						  "2":{"id":2,"name":"Teste 2","quantity":2,"codeValue":"T2","isPublished":false,"expiration":"0001-01-01T00:00:00Z","price":2}
						 }`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, expectedHeader, res.Header())
	})
}

func TestProductsGetById(t *testing.T) {
	t.Run("success to get a product", func(t *testing.T) {
		// given
		hd := setupHandler()

		// when
		req := httptest.NewRequest("GET", "/products/1", nil)
		req = testutils.WithUrlParam(t, req, "productId", "1")
		res := httptest.NewRecorder()
		hd.GetByID()(res, req)

		// then
		expectedCode := http.StatusFound
		expectedBody := `{"id":1,"name":"Teste 1","quantity":1,"codeValue":"T1","isPublished":true,"expiration":"0001-01-01T00:00:00Z","price":1}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, expectedHeader, res.Header())
	})

	t.Run("bad request while trying to get a product", func(t *testing.T) {
		// given
		hd := setupHandler()

		// when
		req := httptest.NewRequest("GET", "/products/invalidID", nil)
		req = testutils.WithUrlParam(t, req, "productId", "invalidId")
		res := httptest.NewRecorder()
		hd.GetByID()(res, req)

		// then
		expectedCode := http.StatusBadRequest
		expectedBody := `{"message":"ID inválido!", "status":"Bad Request"}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, expectedHeader, res.Header())
	})

	t.Run("not found while trying to get a product", func(t *testing.T) {
		// given
		hd := setupHandler()
		// when
		req := httptest.NewRequest("GET", "/products/4", nil)
		req = testutils.WithUrlParam(t, req, "productId", "4")
		res := httptest.NewRecorder()
		hd.GetByID()(res, req)

		// then
		expectedCode := http.StatusNotFound
		expectedBody := `{"message":"Produto não encontrado!", "status":"Not Found"}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, expectedHeader, res.Header())
	})
}

func TestProductsCreate(t *testing.T) {
	t.Run("success to create a product", func(t *testing.T) {
		// given
		hd := setupHandler()
		newProduct := `{"name":"Teste 3","quantity":3,"codeValue":"T3","isPublished":true,"expiration":"11/05/2025","price":3.00}`

		// when
		req := httptest.NewRequest("POST", "/products", strings.NewReader(newProduct))
		res := httptest.NewRecorder()
		// req.Header.Set("Authorization", os.Getenv("API_TOKEN"))

		hd.Create()(res, req)

		// then
		expectedCode := http.StatusCreated
		expectedBody := `{"id":3,"name":"Teste 3","quantity":3,"codeValue":"T3","isPublished":true,"expiration":"2025-05-11T00:00:00Z","price":3}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, expectedHeader, res.Header())
	})

	// t.Run("unauthorized while trying to create a product", func(t *testing.T) {
	// 	// given
	// 	hd := setupHandler()
	// 	newProduct := `{"name":"Teste 3","quantity":3,"codeValue":"T3","isPublished":true,"expiration":"11/05/2025","price":3.00}`

	// 	// when
	// 	req := httptest.NewRequest("POST", "/products", strings.NewReader(newProduct))
	// 	res := httptest.NewRecorder()
	// 	getAllProducts := hd.Create()
	// 	getAllProducts(res, req)

	// 	// then
	// 	expectedCode := http.StatusUnauthorized
	// 	expectedBody := `{"message":"Token inválido.", "status":"Unauthorized"}`
	// 	expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

	// 	require.Equal(t, expectedCode, res.Code)
	// 	require.JSONEq(t, expectedBody, res.Body.String())
	// 	require.Equal(t, expectedHeader, res.Header())
	// })
}

func TestProductsUpdateOrCreate(t *testing.T) {
	t.Run("success to update a product", func(t *testing.T) {
		// given
		hd := setupHandler()
		updatedProduct := `{"id":1,"name":"Teste 3","quantity":2,"codeValue":"T7","isPublished":false,"expiration":"11/05/2027","price":45.00}`

		// when
		req := httptest.NewRequest("PUT", "/products/1", strings.NewReader(updatedProduct))
		req = testutils.WithUrlParam(t, req, "productId", "1")
		// req.Header.Set("Authorization", os.Getenv("API_TOKEN"))
		res := httptest.NewRecorder()
		hd.UpdateOrCreate()(res, req)

		// then
		expectedCode := http.StatusOK
		expectedBody := `{"id":1,"name":"Teste 3","quantity":2,"codeValue":"T7","isPublished":false,"expiration":"2027-05-11T00:00:00Z","price":45.00}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, expectedHeader, res.Header())
	})

	t.Run("bad request while trying to update a product", func(t *testing.T) {
		// given
		hd := setupHandler()

		updatedProduct := `{"id":1,"name":"Teste 3","quantity":2,"codeValue":"T7","isPublished":false,"expiration":"11/05/2027","price":45.00}`

		// when
		req := httptest.NewRequest("PUT", "/products/invalidID", strings.NewReader(updatedProduct))
		req = testutils.WithUrlParam(t, req, "productId", "invalidId")
		// req.Header.Set("Authorization", os.Getenv("API_TOKEN"))
		res := httptest.NewRecorder()
		hd.UpdateOrCreate()(res, req)

		// then
		expectedCode := http.StatusBadRequest
		expectedBody := `{"message":"ID inválido!", "status":"Bad Request"}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, expectedHeader, res.Header())
	})

	// t.Run("unauthorized while trying to update a product", func(t *testing.T) {
	// 	// given
	// 	hd := setupHandler()
	// 	r := chi.NewRouter()
	// 	r.Put("/products/{productId}", hd.UpdateOrCreate())

	// 	updatedProduct := `{"id":1,"name":"Teste 3","quantity":2,"codeValue":"T7","isPublished":false,"expiration":"11/05/2027","price":45.00}`

	// 	// when
	// 	req := httptest.NewRequest("PUT", "/products/1", strings.NewReader(updatedProduct))
	// 	res := httptest.NewRecorder()
	// 	r.ServeHTTP(res, req)

	// 	// then
	// 	expectedCode := http.StatusUnauthorized
	// 	expectedBody := `{"message":"Token inválido.", "status":"Unauthorized"}`
	// 	expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

	// 	require.Equal(t, expectedCode, res.Code)
	// 	require.JSONEq(t, expectedBody, res.Body.String())
	// 	require.Equal(t, expectedHeader, res.Header())
	// })
}

func TestProductsPatch(t *testing.T) {
	t.Run("success to update an item in product", func(t *testing.T) {
		// given
		hd := setupHandler()
		updatedProduct := `{"quantity":2}`

		// when
		req := httptest.NewRequest("PATCH", "/products/1", strings.NewReader(updatedProduct))
		req = testutils.WithUrlParam(t, req, "productId", "1")
		// req.Header.Set("Authorization", os.Getenv("API_TOKEN"))
		res := httptest.NewRecorder()
		hd.Patch()(res, req)

		// then
		expectedCode := http.StatusOK
		expectedBody := `{"id":1,"name":"Teste 1","quantity":2,"codeValue":"T1","isPublished":true,"expiration":"0001-01-01T00:00:00Z","price":1.00}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, expectedHeader, res.Header())
	})

	t.Run("bad request while trying to an item in product", func(t *testing.T) {
		// given
		hd := setupHandler()
		updatedProduct := `{"quantity":2}`

		// when
		req := httptest.NewRequest("PATCH", "/products/invalidId", strings.NewReader(updatedProduct))
		req = testutils.WithUrlParam(t, req, "productId", "invalidId")
		// req.Header.Set("Authorization", os.Getenv("API_TOKEN"))
		res := httptest.NewRecorder()
		hd.Patch()(res, req)

		// then
		expectedCode := http.StatusBadRequest
		expectedBody := `{"message":"ID inválido!", "status":"Bad Request"}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, expectedHeader, res.Header())
	})

	t.Run("not found while trying to an item in product", func(t *testing.T) {
		// given
		hd := setupHandler()
		updatedProduct := `{"quantity":2}`

		// when
		req := httptest.NewRequest("PATCH", "/products/4", strings.NewReader(updatedProduct))
		req = testutils.WithUrlParam(t, req, "productId", "4")
		// req.Header.Set("Authorization", os.Getenv("API_TOKEN"))
		res := httptest.NewRecorder()
		hd.Patch()(res, req)

		// then
		expectedCode := http.StatusNotFound
		expectedBody := `{"message":"Produto não encontrado!", "status":"Not Found"}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, expectedHeader, res.Header())
	})

	// t.Run("unauthorized while trying to update an item in product", func(t *testing.T) {
	// 	// given
	// 	hd := setupHandler()
	// 	r := chi.NewRouter()
	// 	r.Patch("/products/{productId}", hd.Patch())

	// 	updatedProduct := `{"quantity":2}`

	// 	// when
	// 	req := httptest.NewRequest("PATCH", "/products/1", strings.NewReader(updatedProduct))
	// 	res := httptest.NewRecorder()
	// 	r.ServeHTTP(res, req)

	// 	// then
	// 	expectedCode := http.StatusUnauthorized
	// 	expectedBody := `{"message":"Token inválido.", "status":"Unauthorized"}`
	// 	expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

	// 	require.Equal(t, expectedCode, res.Code)
	// 	require.JSONEq(t, expectedBody, res.Body.String())
	// 	require.Equal(t, expectedHeader, res.Header())
	// })
}

func TestProductsDelete(t *testing.T) {
	t.Run("success to delete a product", func(t *testing.T) {
		// given
		hd := setupHandler()

		// when
		req := httptest.NewRequest("DELETE", "/products/1", nil)
		req = testutils.WithUrlParam(t, req, "productId", "1")
		// req.Header.Set("Authorization", os.Getenv("API_TOKEN"))
		res := httptest.NewRecorder()
		hd.Delete()(res, req)

		// then
		expectedCode := http.StatusNoContent
		expectedHeader := http.Header{}

		require.Equal(t, expectedCode, res.Code)
		require.Empty(t, res.Body.String())
		require.Equal(t, expectedHeader, res.Header())
	})

	t.Run("bad request while trying to delete a product", func(t *testing.T) {
		// given
		hd := setupHandler()

		// when
		req := httptest.NewRequest("DELETE", "/products/invalidId", nil)
		req = testutils.WithUrlParam(t, req, "productId", "invalidId")
		// req.Header.Set("Authorization", os.Getenv("API_TOKEN"))
		res := httptest.NewRecorder()
		hd.Delete()(res, req)

		// then
		expectedCode := http.StatusBadRequest
		expectedBody := `{"message":"ID inválido!", "status":"Bad Request"}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, expectedHeader, res.Header())
	})

	t.Run("not found while trying to delete a product", func(t *testing.T) {
		// given
		hd := setupHandler()

		// when
		req := httptest.NewRequest("DELETE", "/products/4", nil)
		req = testutils.WithUrlParam(t, req, "productId", "4")
		// req.Header.Set("Authorization", os.Getenv("API_TOKEN"))
		res := httptest.NewRecorder()
		hd.Delete()(res, req)

		// then
		expectedCode := http.StatusNotFound
		expectedBody := `{"status":"Not Found","message":"Produto não encontrado!"}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, expectedHeader, res.Header())
	})

	// t.Run("unauthorized while trying to delete a product", func(t *testing.T) {
	// 	// given
	// 	hd := setupHandler()
	// 	r := chi.NewRouter()
	// 	r.Delete("/products/{productId}", hd.Delete())

	// 	// when
	// 	req := httptest.NewRequest("DELETE", "/products/4", nil)
	// 	res := httptest.NewRecorder()
	// 	r.ServeHTTP(res, req)

	// 	// then
	// 	expectedCode := http.StatusUnauthorized
	// 	expectedBody := `{"message":"Token inválido.", "status":"Unauthorized"}`
	// 	expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

	// 	require.Equal(t, expectedCode, res.Code)
	// 	require.JSONEq(t, expectedBody, res.Body.String())
	// 	require.Equal(t, expectedHeader, res.Header())
	// })
}
