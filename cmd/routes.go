package cmd

import (
	"net/http"
	"go.mod/handlers"
	"go.mod/middleware"
)

func initRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	//Get All Products
	mux.Handle(
		"GET /products",
		manager.With(
			http.HandlerFunc(
				handlers.GetProducts,
			),
		))

	//Add A Product
	mux.Handle(
		"POST /products",
		manager.With(
			http.HandlerFunc(
				handlers.CreateProduct,
			),
		))

	//Get A Product
	mux.Handle(
		"GET /products/{id}",
		manager.With(
			http.HandlerFunc(
				handlers.GetProductById,
			),
		))
}
