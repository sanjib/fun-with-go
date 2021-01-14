package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

// WordCount provides count of words in provided string.
func WordCount(s string) map[string]int {
	count := map[string]int{}
	words := strings.Fields(s)
	for _, word := range words {
		count[word]++
	}
	return count
}

func main() {
	wc.Test(WordCount)
}
