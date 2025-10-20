package cmd

import (
	"fmt"
	"os"

	"go.mod/config"
	"go.mod/infra/db"
	"go.mod/repo"
	"go.mod/rest"
	"go.mod/rest/handlers/product"
	"go.mod/rest/handlers/user"
	middleware "go.mod/rest/middlewares"
)

func Serve() {
	cnf := config.GetConfig()
	dbCon, err := db.NewConnection(cnf.DB)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	productRepo := repo.NewProductRepo(*dbCon)
	userRepo := repo.NewUserRepo(dbCon)

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
