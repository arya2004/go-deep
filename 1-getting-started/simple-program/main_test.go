package main

import "testing"

func TestGreet(t *testing.T) {
    tests := []struct {
        names []string
        want  string
    }{
        {[]string{}, "Hello, World!"},
        {[]string{"Alice"}, "Hello, Alice!"},
        {[]string{"Alice", "Bob"}, "Hello, Alice, Bob!"},
    }

    for _, tt := range tests {
        got := greet(tt.names)
        if got != tt.want {
            t.Errorf("greet(%v) = %v; want %v", tt.names, got, tt.want)
        }
    }
}
