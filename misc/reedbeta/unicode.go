// Ref article: http://reedbeta.com/blog/programmers-intro-to-unicode/
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println('\u0041')
	fmt.Println(string('\u0041'))
	fmt.Println(utf8.DecodeRuneInString("A"))
	fmt.Println(string('\u03b8'))
}
