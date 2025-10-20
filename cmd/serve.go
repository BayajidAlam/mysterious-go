package cmd

import (
	"go.mod/config"
	"go.mod/repo"
	"go.mod/rest"
	"go.mod/rest/handlers/product"
	"go.mod/rest/handlers/user"
	middleware "go.mod/rest/middlewares"
)

func Serve() {
	cnf := config.GetConfig()

	productRepo := repo.NewProductRepo()
	userRepo := repo.NewUserRepo()

	middlewares := middleware.NewMiddleware(cnf)

	productHandler := product.NewHandler(middlewares, productRepo)
	userHandler := user.NewHandler(cnf, userRepo)

	server := rest.NewServer(
		cnf,
		productHandler,
		userHandler,
	)
	server.Start()
}
