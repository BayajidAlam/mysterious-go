package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func main() {
	result := add(5, 3)
	fmt.Println(result) // Output: 8
}


//HOC
package main

import "fmt"

// higher-order function
func apply(fn func(int) int, val int) int {
	return fn(val)
}

// normal function
func square(x int) int {
	return x * x
}

func main() {
	result := apply(square, 5)
	fmt.Println(result) // Output: 25
}


//Exp 2
package main

import "fmt"

func multiplier(factor int) func(int) int {
	return func(n int) int {
		return n * factor
	}
}

func main() {
	double := multiplier(2)
	triple := multiplier(3)

	fmt.Println(double(5)) // Output: 10
	fmt.Println(triple(5)) // Output: 15
}
