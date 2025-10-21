package product

import (
	"fmt"
	"net/http"
	"strconv"

	"go.mod/utils"
)

func (h *Handler) GetProductById(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")

	pID, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "Please provide a valid ID", 400)
		return
	}

	product, err := h.svc.Get(pID)
	if err != nil {
		fmt.Println(err)
		http.Error(
			w,
			"Internal server error",
			http.StatusInternalServerError,
		)
		return
	}
	if product == nil {
		utils.SendError(w, 404, "Product not found!")
		return
	}
	
	utils.SendData(w, product, 200)
}
