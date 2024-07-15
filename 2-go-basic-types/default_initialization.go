package main

import (
	"fmt"
)

func default_initialization() {
    var i int
    var f float64
    var b bool
    var s string

    fmt.Printf("%v %v %v %q\n", i, f, b, s)
    // Output: 0 0 false ""
}
