package main

import (
	"fmt"
	"strings"
)

func main() {
	text := `
	Needles and pins
	Needles and pins
	Sew me a sail
	To catch me the wind`

	words := strings.Fields(text)

	count := make(map[string]int)
	for _, word := range words {
		//w := strings.ToLower(strings.TrimSpace(word))
		//if _, ok := count[w]; ok {
		//	count[w]++
		//} else {
		//	count[w] = 1
		//}
		count[strings.ToLower(word)]++
	}
	fmt.Println(count)

}
