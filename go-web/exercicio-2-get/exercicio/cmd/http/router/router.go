package router

import (
	"exercicio/cmd/http/middleware"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type router struct {
}

func (router *router) MapRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Auth)

	r.Route("/products", func(rp chi.Router) {
		rp.Mount("/", buildProductsRoutes())
	})

	return r
}

func NewRouter() *router {
	return &router{}
}
