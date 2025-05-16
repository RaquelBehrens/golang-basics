package testutils

import (
	"context"
	"net/http"
	"testing"

	"github.com/go-chi/chi/v5"
)

func WithUrlParam(t *testing.T, r *http.Request, key, value string) *http.Request {
	t.Helper()
	chiCtx := chi.NewRouteContext()
	req := r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, chiCtx))
	chiCtx.URLParams.Add(key, value)
	return req
}

func WithUrlParamst(t *testing.T, r *http.Request, params map[string]string) *http.Request {
	t.Helper()
	chiCtx := chi.NewRouteContext()
	req := r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, chiCtx))
	for key, value := range params {
		chiCtx.URLParams.Add(key, value)
	}
	return req
}

func WithQueryParam(t *testing.T, r *http.Request, key, value string) *http.Request {
	t.Helper()
	q := r.URL.Query()
	q.Add(key, value)
	r.URL.RawQuery = q.Encode()
	return r
}

func WithQueryParams(t *testing.T, r *http.Request, params map[string]string) *http.Request {
	t.Helper()
	q := r.URL.Query()
	for key, value := range params {
		q.Add(key, value)
	}
	r.URL.RawQuery = q.Encode()
	return r
}
