package main

import "fmt"

func main() {
    // ===== Basic Types =====
    var name string = "Bayajid"   // string
    var age int = 25              // int
    height := 5.9                 // float64 (type inferred)
    var isStudent bool = true     // bool
    var grade rune = 'A'          // rune (alias for int32)
    var b byte = 255              // byte (alias for uint8)

    // ===== Arrays =====
    var arr [3]int = [3]int{1, 2, 3}

    // ===== Slices =====
    nums := []int{10, 20, 30, 40}

    // ===== Maps =====
    person := map[string]int{"Alice": 30, "Bob": 25}

    // ===== Struct =====
    type Person struct {
        Name string
        Age  int
    }
    p := Person{Name: "Charlie", Age: 28}

    // ===== Pointer =====
    x := 42
    ptr := &x // pointer to x

    // ===== Empty Interface =====
    var any interface{}
    any = "can hold anything"
    any = 99 // now it's an int

    // ===== Printing Values =====
    fmt.Println("=== Basic Types ===")
    fmt.Println("Name:", name)
    fmt.Println("Age:", age)
    fmt.Println("Height:", height)
    fmt.Println("Is Student:", isStudent)
    fmt.Println("Grade (rune):", string(grade))
    fmt.Println("Byte:", b)

    fmt.Println("\n=== Array & Slice ===")
    fmt.Println("Array:", arr)
    fmt.Println("Slice:", nums)

    fmt.Println("\n=== Map ===")
    fmt.Println("Person map:", person)

    fmt.Println("\n=== Struct ===")
    fmt.Println("Struct:", p)

    fmt.Println("\n=== Pointer ===")
    fmt.Println("Value of x:", x, " | Pointer to x:", ptr, " | Deref:", *ptr)

    fmt.Println("\n=== Empty Interface ===")
    fmt.Println("Any:", any)
}
