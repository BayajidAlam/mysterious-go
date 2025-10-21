package product

import (
	"fmt"
	"net/http"
	"strconv"

	"go.mod/utils"
)

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")

	pID, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "Please give me a valid product id", http.StatusBadRequest)
		return
	}

	err = h.svc.Delete(pID)
	if err != nil {
		fmt.Println(err)
		http.Error(
			w,
			"Internal server error",
			http.StatusInternalServerError,
		)
		return
	}
	
	utils.SendData(
		w,
		"Product deleted successfully",
		200,
	)
}
