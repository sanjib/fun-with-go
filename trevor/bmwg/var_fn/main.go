package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world in trevor")

	whatToSay := "foobar world"
	i := 10
	f := 33.6
	fmt.Printf("wts %s - %d - %f\n", whatToSay, i, f)

	fmt.Println(saySomething())
}

func saySomething() (string, string) {
	s1 := "abc"
	return fmt.Sprintf("Some vals: %s \n", s1), "xxx"
}
