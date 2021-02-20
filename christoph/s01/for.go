package main

import "fmt"

func main() {
	s := []int{1, 2, 3}
	fmt.Println(s)
	for i, _ := range s {
		s[i] = 4
	}
	fmt.Println(s)

	for x := range s {
		fmt.Println(x)
	}

	hello := "hello"
	for x, v := range hello {
		fmt.Println(x, v, string(v))
	}

	word := "abbdf"
	freq := map[string]int{}
	for _, c := range word {
		freq[string(c)]++
	}
	fmt.Println(freq)

	s = []int{1}
	for i, v := range s {
		fmt.Println("i v:", i, v)
	}
}
