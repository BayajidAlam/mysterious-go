package main

import "fmt"

func main() {
    a := 42
    p := &a

    fmt.Println("a =", a)   
    fmt.Println("p =", p)   
    fmt.Println("*p =", *p) 

    *p = 100
    fmt.Println("a after *p change =", a) 
}
