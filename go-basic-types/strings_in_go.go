package main

import (
	"fmt"
	"strings"
)

func strings_in_go() {
    message := "Hello, World!"
    fmt.Println(strings.ToUpper(message)) // "HELLO, WORLD!"
    fmt.Println(strings.Contains(message, "World")) // true
}
