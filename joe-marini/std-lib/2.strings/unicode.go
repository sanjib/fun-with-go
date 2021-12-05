package main

import (
	"log"
	"unicode"
)

func main() {
	const s = "The 'quick' brown fox, jumped over the *LAZY* dog!"

	puncCount := 0
	lowerCount, upperCount := 0, 0
	spaceCount := 0
	hexCount := 0

	for _, r := range s {
		if unicode.IsPunct(r) {
			puncCount++
		}
		if unicode.IsLower(r) {
			lowerCount++
		}
		if unicode.IsUpper(r) {
			upperCount++
		}
		if unicode.IsSpace(r) {
			spaceCount++
		}
		if unicode.Is(unicode.Hex_Digit, r) {
			hexCount++
		}
	}

	log.Println(puncCount)
	log.Println(lowerCount)
	log.Println(upperCount)
	log.Println(spaceCount)
	log.Println(hexCount)
}
