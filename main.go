package main

import (
	"fmt"

	"go.mod/cmd"
	"go.mod/utils"
)

func main() {
	

	jwt, err := utils.CreateJwt("my-secret", utils.Payload{
		Sub:         34,
		FirstName:   "Bayajid",
		LastName:    "Alam",
		Email:       "Bayajid@gmail.com",
		IsShopOwner: false,
	})
	if err != nil {
		fmt.Println("Error occured", err)
	}
	fmt.Println(jwt)

	cmd.Serve()
}
