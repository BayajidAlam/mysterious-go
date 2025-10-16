package handlers

import (
	"encoding/json"
	"fmt"
	"go.mod/database"
	"go.mod/utils"
	"net/http"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct database.Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please give me a valid JSON", 400)
		return
	}

	newProduct.ID = len(database.ProductList) + 1
	database.ProductList = append(database.ProductList, newProduct)

	utils.SendData(w, database.ProductList, 201)
}
