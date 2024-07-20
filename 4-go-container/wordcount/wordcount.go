package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
    // Initialize a map to count words
    wordCount := make(map[string]int)
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanWords)

    for scanner.Scan() {
        word := strings.ToLower(scanner.Text())
        wordCount[word]++
    }

    // Convert map to a slice of key-value pairs
    type kv struct {
        Key   string
        Value int
    }

    var ss []kv
    for k, v := range wordCount {
        ss = append(ss, kv{k, v})
    }

    // Sort slice by values in descending order
    sort.Slice(ss, func(i, j int) bool {
        return ss[i].Value > ss[j].Value
    })

    // Print the top 3 words
    for i, kv := range ss {
        if i >= 3 {
            break
        }
        fmt.Printf("%s: %d\n", kv.Key, kv.Value)
    }
}
