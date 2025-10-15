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

	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Please give me GET request", 400)
		return
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(ProductList)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Origin", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Please give me POST request", 400)
		return
	}

	var newProduct Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please give me a valid JSON", 400)
		return
	}

	newProduct.ID = len(ProductList) + 1
	ProductList = append(ProductList, newProduct)

	w.WriteHeader(201)
	encoder := json.NewEncoder(w)
	encoder.Encode(ProductList)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/products", getProducts)
	mux.HandleFunc("/products/create", createProduct)

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
