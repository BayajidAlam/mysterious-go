package cmd

import (
	"fmt"
	"go.mod/global_router"
	"go.mod/handlers"
	"net/http"
)

func Serve() {
	mux := http.NewServeMux()

	mux.Handle("GET /products", http.HandlerFunc(handlers.GetProducts))
	mux.Handle("POST /products", http.HandlerFunc(handlers.CreateProduct))
	mux.Handle("GET /products/{id}", http.HandlerFunc(handlers.GetProductById))

	fmt.Println("Server is running on port 3000")

	globalRouter := global_router.GlobalRouter(mux)
	err := http.ListenAndServe(":3000", globalRouter)

	if err != nil {
		fmt.Println("Error starting the server: ", err)
	}
}
