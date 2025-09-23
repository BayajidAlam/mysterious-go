package main

import "fmt"

func sum(num1 int, num2 int) {
	sum := num1 + num2
	fmt.Println(sum)
}

func main() {
	a := 10
	b := 20
	sum(a, b)
}
