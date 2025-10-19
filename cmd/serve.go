package cmd

import (
	"go.mod/config"
	"go.mod/rest"
	"go.mod/rest/handlers/product"
	"go.mod/rest/handlers/user"
	middleware "go.mod/rest/middlewares"
)

func Serve() {
	cnf := config.GetConfig()

	// create middleware instance using the constructor
	middlewares := middleware.NewMiddleware(cnf)

	productHandler := product.NewHandler(middlewares)
	userHandler := user.NewHandler()

	server := rest.NewServer(
		cnf,
		productHandler,
		userHandler,
	)
	server.Start()
}
