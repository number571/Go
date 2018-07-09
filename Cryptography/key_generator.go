package main

import (
    "fmt"
    "crypto/rand"
)

func main() {
    fmt.Println(string(generate(30)))
}

func generate(max int16) []byte {
    var slice []byte = make([]byte, max)
    _, err := rand.Read(slice)
    if err != nil { return nil }
    for max = max - 1; max >= 0; max -- {
        slice[max] = slice[max] % 95 + 33
    }
    return slice
}
