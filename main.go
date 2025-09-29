package main

import "fmt"

// Package-level (global) variable
var globalVar = "আমি global!"

func main() {
    fmt.Println("Global:", globalVar) // Access from main

    // Function-level (local) variable
    localVar := "আমি local!"
    fmt.Println("Local:", localVar)

    if true {
        // Block-level variable
        blockVar := "আমি block-level!"
        fmt.Println("Inside if block:", blockVar)
    }
    // fmt.Println(blockVar) // ❌ Error: undefined: blockVar

    // Loop-level variable
    for i := 0; i < 3; i++ {
        loopVar := i * 2
        fmt.Println("LoopVar:", loopVar)
    }
    // fmt.Println(loopVar) // ❌ Error: undefined: loopVar
}
