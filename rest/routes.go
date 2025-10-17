package rest

import (
	"net/http"

	"go.mod/rest/handlers"
	middleware "go.mod/rest/middlewares"
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

	//Update A Product
	mux.Handle(
		"PUT /products/{id}",
		manager.With(
			http.HandlerFunc(
				handlers.UpdateProduct,
			),
		))

	//Delete A Product
	mux.Handle(
		"DELETE /products/{id}",
		manager.With(
			http.HandlerFunc(
				handlers.DeleteProduct,
			),
		))

	//Add A Product
	mux.Handle(
		"POST /users",
		manager.With(
			http.HandlerFunc(
				handlers.CreateUser,
			),
		))
}
