package main

import "fmt"

func changeSlice(s []int) {
	s[0] = 7
}

func main() {
	s1 := []int{1}
	fmt.Println("s1 before changeSlice", s1)
	changeSlice(s1)
	fmt.Println("s1 after changeSlice", s1)
}
