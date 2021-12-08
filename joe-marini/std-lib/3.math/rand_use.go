package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	// shuffle
	rand.Seed(time.Now().UnixNano())
	const numString = "one two three four five six"
	words := strings.Fields(numString)
	rand.Shuffle(len(words), func(i, j int) {
		words[i], words[j] = words[j], words[i]
	})
	fmt.Println("words", words)

	// random string
	const upper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const lower = "abcdefghijklmnopqrstuvwxyz"
	const digits = "0123456789"
	const special = "~!@#$%^&*()_+-={}|?"
	allchars1 := strings.Split(upper+lower+digits+special, "")
	allchars2 := upper + lower + digits + special
	var sb1 strings.Builder
	//var sb2 strings.Builder
	length := 10

	for i := 0; i < length; i++ {
		index := rand.Intn(len(allchars1) - 1)
		sb1.WriteString(allchars1[index])
	}
	fmt.Println("sb1", sb1.String())

	buf := make([]byte, length)
	buf[0] = digits[rand.Intn(len(digits))]
	buf[1] = special[rand.Intn(len(special))]
	buf[2] = upper[rand.Intn(len(upper))]
	buf[3] = lower[rand.Intn(len(lower))]
	for i := 4; i < length; i++ {
		index := rand.Intn(len(allchars1) - 1)
		buf[i] = allchars2[index]
	}
	rand.Shuffle(len(buf), func(i, j int) {
		buf[i], buf[j] = buf[j], buf[i]
	})
	fmt.Println("sb2", string(buf))
}
