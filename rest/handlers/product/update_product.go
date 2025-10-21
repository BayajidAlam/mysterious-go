package product

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"go.mod/domain"
	"go.mod/utils"
)

type RequestUpdateProduct struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageUrl    string `json:"imageUrl"`
}

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")
	pID, err := strconv.Atoi(productID)
	if err != nil {
		utils.SendError(
			w,
			http.StatusBadRequest,
			"Invalid product ID",
		)
		return
	}

	var req RequestUpdateProduct
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&req)
	if err != nil {
		utils.SendError(
			w,
			http.StatusBadRequest,
			"Invalid request body",
		)
		return
	}

	_, err = h.svc.Update(domain.Product{
		ID:          pID,
		Title:       req.Title,
		Description: req.Description,
		ImageUrl:    req.ImageUrl,
	})
	if err != nil {
		fmt.Println(err)
		http.Error(
			w,
			"Internal server error",
			http.StatusInternalServerError,
		)
		return
	}

	utils.SendData(w, "Product Updated", http.StatusOK)
}
