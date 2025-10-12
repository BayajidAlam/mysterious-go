package main

import "fmt"

func a() {
	i := 0

	defer fmt.Println(i)

	i = i + 1

	return
}

func main() {
	a()

	for i := 0; i < 3; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}

}
