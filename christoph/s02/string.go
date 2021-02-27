package main

import "fmt"

func main() {
	s1 := "aeo"
	fmt.Println("aeo", len(s1))

	s2 := "áéö"
	fmt.Println("áéö", len(s2))

	s3 := "hello"
	b30 := s3[0]
	fmt.Println("b30 as byte", b30)
	fmt.Println("b30 as string", string(b30))
}
