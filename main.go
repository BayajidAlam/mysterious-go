package main

// 2.1 Basic Function
func add(a int, b int) int {
    return a + b
}

//2.2 Multiple Return Values
func divide(a, b int) (int, int) {
    return a / b, a % b
}

//2.3 Named Return Values
func getValues() (x int, y int) {
    x = 10
    y = 20
    return
}

//2.4 Variadic Functions
func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}

//2.5 Anonymous Functions

multiply := func(a, b int) int {
    return a * b
}

//2.6 Higher-Order Functions
func apply(f func(int, int) int, a int, b int) int {
    return f(a, b)
}

//2.7 Recursive Functions
func factorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * factorial(n-1)
}

//2.8 Methods (Function with Receiver)
type Circle struct {
    radius float64
}

func (c Circle) Area() float64 {
    return 3.14 * c.radius * c.radius
}

//2.9 Closures
counter := 0
increment := func() int {
    counter++
    return counter
}

//2.10 Deferred Functions
defer fmt.Println("World")
fmt.Println("Hello")
// Output:
// Hello
// World