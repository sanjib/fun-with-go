// Counts the characters in a string,
// ignores uppercase characters.
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, c := range os.Args[1] {
		s := strings.ToLower(string(c))
		counts[s]++
	}
	fmt.Println(counts)
}