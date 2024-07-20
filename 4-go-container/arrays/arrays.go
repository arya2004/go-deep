package main

import "fmt"

func main() {
    // Fixed-size buffer example
    var buffer [10]int // fixed-size buffer of 10 elements
    for i := 0; i < 10; i++ {
        buffer[i] = i * i
    }
    fmt.Println(buffer) // Output: [0 1 4 9 16 25 36 49 64 81]

    // Array copying example
    a := [3]int{1, 2, 3}
    b := a                   // b is a copy of a
    b[0] = 7                 // modifying b doesn't affect a
    fmt.Println(a, b)        // Output: [1 2 3] [7 2 3]

    // Different types example
    var m [4]int
    var c [3]int
    // m = c // This will not compile because they are different types

	
}
