package product

import (
	"net/http"
	"strconv"

	"go.mod/database"
	"go.mod/utils"
)

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")

	pID, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "Please give me a valid product id", http.StatusBadRequest)
		return
	}

	database.Delete(pID)

	utils.SendData(
		w,
		"Product deleted successfully",
		200,
	)
}
