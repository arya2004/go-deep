package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
    fmt.Println(greet(os.Args[1:]))
}

func greet(names []string) string {
    if len(names) == 0 {
        return "Hello, World!"
    }
    return "Hello, " + strings.Join(names, ", ") + "!"
}
