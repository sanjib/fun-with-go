package main

import "fmt"

func main() {
	a := [8]int{}
	fmt.Println("a", a)

	s := a[3:7]
	fmt.Println("s", s)

	fmt.Println("len s", len(s)) // 4 = 7 -3
	fmt.Println("cap s", cap(s)) // 5 = len(a) - 3
	fmt.Println("--")

	s1 := a[4:7]
	s2 := a[3:6]

	s1[0] = 42
	fmt.Println("s1", s1) // s1[0] = 42
	fmt.Println("s2", s2) // s2[1] = 42
	fmt.Println("--")

	var sn1 []int
	fmt.Printf("len sn1 %d, %#v\n", len(sn1), sn1)
	sn2 := []int{}
	fmt.Printf("len sn1 %d, %#v\n", len(sn2), sn2)
}
