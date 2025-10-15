package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageUrl    string `json:"imageUrl"`
}

var ProductList []Product

func getProducts(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		http.Error(w, "Please give me GET request", 400)
		return
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(ProductList)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/products", getProducts)

	fmt.Println("Server is running on port 3000")
	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		fmt.Println("Error starting the server: ", err)
	}
}

func init() {
	ProductList = append(ProductList,
		Product{
			ID:          1,
			Title:       "Laptop",
			Description: "A high performance laptop",
			ImageUrl:    "https://example.com/laptop.jpg",
		},
		Product{
			ID:          2,
			Title:       "Smartphone",
			Description: "Latest model smartphone",
			ImageUrl:    "https://example.com/smartphone.jpg",
		},
		Product{
			ID:          3,
			Title:       "Headphones",
			Description: "Noise cancelling headphones",
			ImageUrl:    "https://example.com/headphones.jpg",
		},
	)
}
