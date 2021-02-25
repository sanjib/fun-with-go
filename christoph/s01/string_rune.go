package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "🐱🐰"
	r, size := utf8.DecodeRuneInString(str)
	fmt.Println("r, size:", r, size)
	fmt.Println(string(r))
	fmt.Println(str[size:])
}
