package main

import (
	"fmt"
	"strings"
)

func main() {
	s2 := "áéö唖钳辮."
	fmt.Println(s2, len(s2))

	for _, c := range s2 {
		fmt.Println(len(string(c)), string(c))
	}

	fmt.Printf("index rune: %#v\n", strings.IndexRune(s2, 'é'))
	fmt.Printf("index byte: %#v\n", strings.IndexByte(s2, 'é'))
}
