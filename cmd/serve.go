package cmd

import (
	"go.mod/config"
	"go.mod/rest"
	"go.mod/rest/handlers/product"
	"go.mod/rest/handlers/user"
)

func Serve() {
	cnf := config.GetConfig()

	productHandler := product.NewHandler()
	userHandler := user.NewHandler()

	server := rest.NewServer(
		cnf,
		productHandler,
		userHandler,
	)
	server.Start()
}
