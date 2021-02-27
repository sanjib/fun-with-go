package main

import "fmt"

type s1 struct {
	n int
	b bool
}

type s2 struct {
	r []rune
}

type s3 struct {
	r [3]rune
}

func main() {
	s11 := s1{
		n: 4,
		b: true,
	}
	s12 := s1{
		n: 4,
		b: true,
	}
	fmt.Println(s11 == s12)

	s21 := s2{r: []rune{'a', 'b', '🎵'}}
	s22 := s2{r: []rune{'a', 'b', '🎶'}}
	// slices can't be compared
	//fmt.Println(s21 == s22)
	fmt.Println(s21, s22)

	s31 := s2{r: []rune{'a', 'b', '🎵'}}
	s32 := s2{r: []rune{'a', 'b', '🎶'}}
	// arrays can't be compared
	//fmt.Println(s31 = s32)
	fmt.Println(s31, s32)
}
