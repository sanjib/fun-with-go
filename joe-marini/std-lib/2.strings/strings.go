package main

import "fmt"

func main() {
	s := "The quick brown fox jumps over the lazy dog"
	fmt.Println(len(s))

	for _, c := range s {
		fmt.Printf("%v,", c)
	}
}
