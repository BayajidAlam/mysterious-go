package handlers

import (
	"encoding/json"
	"fmt"
	"go.mod/database"
	"go.mod/utils"
	"net/http"
)

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var reqLogin ReqLogin
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&reqLogin)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid Request Data",
			http.StatusBadRequest,
		)
		return
	}

	usr := database.Find(
		reqLogin.Email,
		reqLogin.Password,
	)

	fmt.Println(usr,"user")

	if usr == nil {
		http.Error(
			w,
			"Invalid credentials",
			http.StatusBadRequest,
		)
		return
	}

	utils.SendData(
		w,
		usr,
		http.StatusCreated,
	)
}
