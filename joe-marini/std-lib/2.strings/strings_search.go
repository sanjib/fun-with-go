package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "The quick brown fox jumps over the lazy dog"
	vowels := "aeiouAEIOU"
	baseFilename1 := "temp_picfile.jpg"
	baseFilename2 := "asdf.txt"

	fmt.Println(strings.Contains(s, "jump")) // true
	fmt.Println(strings.Contains(s, "abcd")) // false
	fmt.Println(strings.Index(s, "fox"))     // 16
	fmt.Println(strings.Index(s, "cat"))     // -1

	fmt.Println(strings.IndexAny(vowels, "beau")) // 0
	fmt.Println(strings.IndexAny("beau", vowels)) // 1

	fmt.Println(strings.HasPrefix(baseFilename2, "temp")) // false
	fmt.Println(strings.HasSuffix(baseFilename1, "txt"))  // false
	fmt.Println(strings.HasPrefix(baseFilename1, "temp")) // true
	fmt.Println(strings.HasSuffix(baseFilename2, "txt"))  // true

	fmt.Println(strings.Count(s, "the")) // 1
	fmt.Println(strings.Count(s, "he"))  // 2
}
