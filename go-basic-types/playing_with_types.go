package main

import (
	"fmt"
)

func playing_with_types() {
    var y int = 10
    b := 15.5

    fmt.Printf("y: %T, %v\n", y, y)
    fmt.Printf("b: %T, %v\n", b, b)

    y = int(b)
    fmt.Printf("y after conversion: %T, %v\n", y, y)
}
