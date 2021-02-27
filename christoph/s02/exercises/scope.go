package main

import (
	"fmt"
	"unicode"
)

func main() {
	word := "abcde"
	for _, char := range word {
		upper := unicode.ToUpper(char)
		fmt.Print(string(upper))
	}
	fmt.Println("\n" + word)
}
