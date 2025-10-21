package product

import (
	"encoding/json"
	"fmt"
	"net/http"

	"go.mod/domain"
	"go.mod/utils"
)

type RequestCreateProduct struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageUrl    string `json:"imageUrl"`
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {

	var req RequestCreateProduct
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		fmt.Println(err)
		http.Error(
			w,
			"Please give me a valid JSON",
			400,
		)
		return
	}

	createdProduct, err := h.svc.Create(domain.Product{
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

	utils.SendData(w, createdProduct, 201)
}
