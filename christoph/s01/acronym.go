// acronym creates acronym from supplied argument to the program. If no argument
//is provided then the default value "Pan Galactic Gargle Blaster" is used.
package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func acronym(s string) (acr string) {
	var acrChars []string
	for _, word := range strings.Split(s, " ") {
		r, _ := utf8.DecodeRuneInString(word)
		acrChars = append(acrChars, string(r))
	}
	acr = strings.Join(acrChars, "")
	return
}

func main() {
	s := "Pan Galactic Gargle Blaster"
	if len(os.Args) > 1 {
		s = strings.Join(os.Args[1:], " ")
	}
	fmt.Println(acronym(s))
}
