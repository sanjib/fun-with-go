package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("abcd", "bc"))
	fmt.Println(strings.HasPrefix("abcd", "abc"))
	fmt.Println(strings.ToLower("ABC"))
	fmt.Println(strings.Trim("http:/abc", "http"))
	fmt.Println(strings.Count("abcdef", "ab"))
	fmt.Println(strings.Index("abcdef", "bc"))
	fmt.Println(strings.Index("abcdef", "bd"))
	fmt.Println(strings.Join([]string{"a", "b", "c"}, "-"))
	fmt.Println(strings.Split("hello", ""))
}
