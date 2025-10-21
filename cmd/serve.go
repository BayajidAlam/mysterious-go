package cmd

import (
	"fmt"
	"os"

	"go.mod/config"
	"go.mod/infra/db"
	"go.mod/product"
	"go.mod/repo"
	"go.mod/rest"
	productHandler "go.mod/rest/handlers/product"
	userHandler "go.mod/rest/handlers/user"
	middleware "go.mod/rest/middlewares"
	"go.mod/user"
)

func Serve() {
	cnf := config.GetConfig()
	dbCon, err := db.NewConnection(cnf.DB)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	db.DbMigrate(dbCon, "./migrations")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//Repos
	productRepo := repo.NewProductRepo(*dbCon)
	userRepo := repo.NewUserRepo(dbCon)

	//Domains
	userSvc := user.NewService(userRepo)
	productSvc := product.NewService(productRepo)

	middlewares := middleware.NewMiddleware(cnf)

	//Handlers
	productHandler := productHandler.NewHandler(middlewares, productSvc)
	userHandler := userHandler.NewHandler(cnf, userSvc)

	server := rest.NewServer(
		cnf,
		productHandler,
		userHandler,
	)
	server.Start()
}
