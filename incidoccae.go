package main

import (
    "fmt"
    "strings"
)

func main() {
    sentence := "Splits it into words."
    words := strings.Split(sentence, " ")
    fmt.Println(words)
    // Output: [Splits it into words.]
}
