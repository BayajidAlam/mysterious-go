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
	
	product := database.Get(pID)
	if product == nil {
		utils.SendError(w, 404, "Product not found!")
		return
	}
	utils.SendData(w, product, 200)
}
