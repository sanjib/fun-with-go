package main

import (
	"fmt"
)

func main() {
	s := "🐭"
	fmt.Println("len", len(s))
	var sBytes []byte
	for i := 0; i < len(s); i++ {
		sBytes = append(sBytes, s[i])
		fmt.Println("i, s[i] string(s[i])", i, s[i], string(s[i]))
	}
	fmt.Println("reconstructing back the mouse", string(sBytes))
}
