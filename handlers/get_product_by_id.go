package handlers

import (
	"net/http"
	"strconv"

	"go.mod/database"
	"go.mod/utils"
)

func GetProductById(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")

	pID, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "Please provide a valid ID", 400)
		return
	}
	for _, product := range database.ProductList {
		if product.ID == pID {
			utils.SendData(w, product, 200)
			return
		}
	}

	utils.SendData(w, "Product not found!", 404)
}
