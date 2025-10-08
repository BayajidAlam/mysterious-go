package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func (user User) printUser() {
	fmt.Println("User name: ", user.Name)
	fmt.Println("User age: ", user.Age)
}

func main() {
	var user1 User

	user1 = User{
		Name: "Bayajid",
		Age:  19,
	}
	user1.printUser()
}
