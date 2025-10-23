package product

import (
	"net/http"
	"strconv"

	"go.mod/utils"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {

	reqQuery := r.URL.Query()
	pageAsStr := reqQuery.Get("page")
	limitAsStr := reqQuery.Get("limit")

	page, _ := strconv.ParseInt(pageAsStr, 10, 32)
	limit, _ := strconv.ParseInt(limitAsStr, 10, 32)

	if page <= 0 {
		page = 1
	}

	if limit <= 0 {
		limit = 10
	}

	products, err := h.svc.List(page, limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cnt, err := h.svc.Count()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.SendPage(w, products, page, limit, cnt)
}
