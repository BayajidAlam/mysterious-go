package main

import "fmt"

var (
	a = 10
	b = 30
)

func printName(num int) {
	fmt.Println(num)
}

func add(a int, b int) {
	res := a + b
	printName(res)
}

func main() {
	add(a, b)
}
