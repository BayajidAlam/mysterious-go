package main

import "fmt"

// Global variable
var globalVar = "I am global"

func main() {
    // Local variable
    localVar := "I am local"

    fmt.Println("Inside main:")
    fmt.Println(globalVar) // accessible
    fmt.Println(localVar)  // accessible

    printSomething()
}

func printSomething() {
    fmt.Println("Inside printSomething function:")
    fmt.Println(globalVar) // accessible
    // fmt.Println(localVar) // ❌ Error: localVar not accessible here
}
