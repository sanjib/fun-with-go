package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// \x80 is an illegal UTF-8 encoding
	for pos, char := range "日本\x80語" {
		str := string(char)
		if char == utf8.RuneError {
			str = "illegal char"
		}
		if char == '\uFFFD' {
			fmt.Println("illegal utf-8 char encountered")
		}
		fmt.Printf("%d - %d - %s\n", pos, char, str)
	}
}
