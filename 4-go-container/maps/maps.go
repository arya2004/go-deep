package main

import (
	"fmt"
	"sort"
)

func main() {
    // Declaring maps
    m := make(map[string]int) // Create an empty map
    n := map[string]int{"foo": 1, "bar": 2} // Map with initial values
    fmt.Println(m, n)

    // Adding and retrieving values
    m["foo"] = 1
    value := m["foo"]
    fmt.Println(value) // Output: 1

    // Checking for key existence
    if value, ok := m["foo"]; ok {
        fmt.Println("Key exists with value", value)
    } else {
        fmt.Println("Key does not exist")
    }

    // Deleting keys
    delete(m, "foo")
    if value, ok := m["foo"]; ok {
        fmt.Println("Key exists with value", value)
    } else {
        fmt.Println("Key does not exist")
    }

    // Iterating over maps
    m = map[string]int{"foo": 1, "bar": 2}
    for k, v := range m {
        fmt.Println(k, v)
    }

    // Practical use of maps
    words := []string{"apple", "banana", "apple", "orange", "banana", "apple"}
    wordCount := make(map[string]int)

    for _, word := range words {
        wordCount[word]++
    }

    fmt.Println(wordCount) // Output: map[apple:3 banana:2 orange:1]

    // Maps and slices together
    type kv struct {
        Key   string
        Value int
    }

    var ss []kv
    for k, v := range wordCount {
        ss = append(ss, kv{k, v})
    }

    sort.Slice(ss, func(i, j int) bool {
        return ss[i].Value > ss[j].Value
    })

    for _, kv := range ss {
        fmt.Printf("%s: %d\n", kv.Key, kv.Value)
    }
}
