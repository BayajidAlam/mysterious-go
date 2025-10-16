package cmd

import (
	"fmt"
	"net/http"

	"go.mod/global_router"
	"go.mod/middleware"
)

func Serve() {
	manager := middleware.NewManager()
	mux := http.NewServeMux()

	manager.Use(
		middleware.Logger,
		middleware.Beauty,
	)
	initRoutes(mux, manager)

	fmt.Println("Server is running on port 3000")

	globalRouter := global_router.GlobalRouter(mux)
	err := http.ListenAndServe(":3000", globalRouter)

	if err != nil {
		fmt.Println("Error starting the server: ", err)
	}
}
