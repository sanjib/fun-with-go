package main

import (
	"log"
	"strings"
)

func main() {
	shift := 4
	s := "The quick brown fox jumps over the lazy dog"

	transform := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			val := int('A') + (int(r) - int('A') + shift)
			if val > 91 {
				val -= 26
			} else if val < 65 {
				val += 26
			}
			return rune(val)
		case r >= 'a' && r <= 'z':
			val := int('a') + (int(r) - int('a') + shift)
			if val > 122 {
				val -= 26
			} else if val < 97 {
				val += 26
			}
			return rune(val)
		}
		return r
	}

	encode := strings.Map(transform, s)
	log.Println(encode)

	shift = -shift
	decode := strings.Map(transform, encode)
	log.Println(decode)

}
