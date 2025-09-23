package main

import "fmt"

func main() {
	a := 2
	if a >= 18 {
		fmt.Println("Get married!")
	} else {
		fmt.Println("Can't get married!")
	}

	day := 3

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Other day")
	}
}
