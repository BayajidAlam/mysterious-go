package product

import (
	"net/http"

	middleware "go.mod/rest/middlewares"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	//Get All Products
	mux.Handle(
		"GET /products",
		manager.With(
			http.HandlerFunc(
				h.GetProducts,
			),
			middleware.AuthenticateJWT,
		))

	//Add A Product
	mux.Handle(
		"POST /products",
		manager.With(
			http.HandlerFunc(
				h.CreateProduct,
			),
		))

	//Get A Product
	mux.Handle(
		"GET /products/{id}",
		manager.With(
			http.HandlerFunc(
				h.GetProductById,
			),
		))

	//Update A Product
	mux.Handle(
		"PUT /products/{id}",
		manager.With(
			http.HandlerFunc(
				h.UpdateProduct,
			),
		))

	//Delete A Product
	mux.Handle(
		"DELETE /products/{id}",
		manager.With(
			http.HandlerFunc(
				h.DeleteProduct,
			),
		))
}
