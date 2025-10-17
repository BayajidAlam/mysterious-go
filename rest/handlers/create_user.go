package handlers

import (
	"encoding/json"
	"fmt"
	"go.mod/database"
	"go.mod/utils"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newProduct database.Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please give me a valid JSON", 400)
		return
	}

	createdProduct := database.Store(newProduct)
	utils.SendData(w, createdProduct, 201)
}
