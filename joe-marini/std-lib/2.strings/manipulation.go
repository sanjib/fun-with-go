package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	s := "the quick brown   fox jumps over the lazy dog"
	s2 := []string{"one", "two", "three", "four"}
	s3 := "This is a string. With some punctionation, for a demo! Yep."

	sub1 := strings.Split(s, " ")
	fmt.Printf("%q\n", sub1)

	sub2 := strings.Split(s, "the")
	fmt.Printf("%q\n", sub2)

	res1 := strings.Join(s2, "-")
	fmt.Printf("%s\n", res1)

	res2 := strings.Fields(s)
	fmt.Printf("%q\n", res2)

	res3 := strings.Fields(s3)
	fmt.Printf("%q\n", res3)

	f := func(r rune) bool {
		return unicode.IsPunct(r)
	}
	res4 := strings.FieldsFunc(s3, f)
	fmt.Printf("%q\n", res4)

	rep := strings.NewReplacer(".", "...", ",", ",,,", "!", "!!!")
	res5 := rep.Replace(s3)
	fmt.Printf("%q\n", res5)
}
