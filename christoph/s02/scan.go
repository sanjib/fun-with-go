package main

import (
	"fmt"
	"log"
)

func main() {
	var s string
	var n int

	fmt.Println("Please enter a word and a number.")
	_, err := fmt.Scanf("a %s%d", &s, &n)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("%#v, %#v", s, n)
}
