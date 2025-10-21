package product

import (
	"net/http"

	"go.mod/utils"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		return
	}

	products, err := h.svc.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.SendData(w, products, 200)
}
