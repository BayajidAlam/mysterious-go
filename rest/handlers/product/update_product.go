package product

import (
	"encoding/json"
	"net/http"
	"strconv"

	"go.mod/database"
	"go.mod/utils"
)

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")
	pID, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "Please give me a valid product id", 400)
		return
	}

	var newProduct database.Product
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&newProduct)
	if err != nil {
		http.Error(w, "Please give me a valid JSON", 400)
		return
	}

	newProduct.ID = pID
	database.Update(newProduct)

	utils.SendData(w, newProduct, 201)
}
