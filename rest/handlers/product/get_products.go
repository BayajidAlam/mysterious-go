package product

import (
	"go.mod/database"
	"go.mod/utils"
	"net/http"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		return
	}

	utils.SendData(w, database.List(), 200)
}
