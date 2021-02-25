package main

import (
	"fmt"
	"strings"
)

func longest(s ...string) int {
	var l int
	for _, word := range s {
		if c := strings.Count(word, ""); c > l {
			l = c
		}
	}
	return l
}

func main() {
	fmt.Println(longest("Six", "sleek", "swans", "swam", "swiftly", "southwards"))
}
