package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println("hello \x61pple")
	rawStr := `
Hello world
  - x
  - y
  - z
Goodbye!
`
	fmt.Println(rawStr)

	mouse := "🐭"
	fmt.Println(utf8.DecodeRuneInString(mouse))

	var mouseRune rune = 128045
	fmt.Println(string(mouseRune))
}
