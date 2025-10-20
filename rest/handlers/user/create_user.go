package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"go.mod/repo"
	"go.mod/utils"
)

type RequestCreateUser struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req RequestCreateUser
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid Request Data",
			http.StatusBadRequest,
		)
		return
	}

	createdUser, err := h.userRepo.Create(repo.User{
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Email:       req.Email,
		Password:    req.Password,
		IsShopOwner: req.IsShopOwner,
	})
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	utils.SendData(
		w,
		createdUser,
		http.StatusCreated,
	)
}
