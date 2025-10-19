package user

import (
	"net/http"

	middleware "go.mod/rest/middlewares"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {

	//Add A User
	mux.Handle(
		"POST /users",
		manager.With(
			http.HandlerFunc(
				h.CreateUser,
			),
		))

	//Login A User
	mux.Handle(
		"POST /login",
		manager.With(
			http.HandlerFunc(
				h.LoginUser,
			),
		))
}
