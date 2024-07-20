package main

import "fmt"

func main() {
    // Declaring slices
    var s []int              // A slice of integers
    s = make([]int, 3)       // Create a slice with length 3
    t := []int{1, 2, 3}      // Slice with initial values
    fmt.Println(s, t)

    // Appending to slices
    s = append(s, 4, 5)      // Now s is [0 0 0 4 5]
    fmt.Println(s)

    // Slicing arrays
    arr := [5]int{1, 2, 3, 4, 5}
    u := arr[1:4]            // u is [2, 3, 4]
    fmt.Println(u)

    // Slicing slices
    v := t[1:3]              // v is [2, 3]
    fmt.Println(v)

    // Copying slices
    a := []int{1, 2, 3}
    b := a                   // b points to the same array as a
    b[0] = 7
    fmt.Println(a, b)        // Output: [7 2 3] [7 2 3]

    // Slice capacity and length
    fmt.Println(len(s), cap(s)) // Output: 5, 6
    s = s[:3]
    fmt.Println(len(s), cap(s)) // Output: 3, 6
    s = append(s, 6)
    fmt.Println(len(s), cap(s)) // Output: 4, 6

    // Practical use of slices
    var list []int
    for i := 0; i < 10; i++ {
        list = append(list, i*i)
    }
    fmt.Println(list) // Output: [0 1 4 9 16 25 36 49 64 81]
}
