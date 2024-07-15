package main

import (
	"fmt"
)

func slices_and_maps() {
    // Slices
    numbers := []int{3, 6, 9, 12, 15}
    fmt.Println(numbers)
    numbers = append(numbers, 18)
    fmt.Println(numbers)

    // Maps
    ages := make(map[string]int)
    ages["Charlie"] = 22
    ages["Dana"] = 28
    fmt.Println(ages)
    fmt.Println("Charlie’s age:", ages["Charlie"])
}
