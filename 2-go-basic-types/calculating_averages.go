package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func calculating_averages() {
    var sum float64
    var count int

    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        val, err := strconv.ParseFloat(scanner.Text(), 64)
        if err != nil {
            fmt.Fprintln(os.Stderr, "Invalid input:", err)
            return
        }
        sum += val
        count++
    }

    if count == 0 {
        fmt.Println("No values provided")
        return
    }

    fmt.Printf("Average: %.2f\n", sum/float64(count))
}
