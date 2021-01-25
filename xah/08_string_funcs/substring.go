package main

import (
	"fmt"
	"strings"
)

func main() {
	var x = "012345"

	fmt.Println(x[1:4])
	fmt.Println(x[2:2])
	fmt.Println(x[3:4])
	fmt.Println("----")

	x = "♥"
	fmt.Println(len(x))
	fmt.Printf("%#v\n", x[:1])
	fmt.Printf("%#v\n", x[1:2])
	fmt.Printf("%#v\n", x[2:])
	fmt.Println("----")

	fmt.Println("\xe2" + "\x99" + "\xa5")
	fmt.Println(strings.Join([]string{"\xe2", "\x99", "\xa5"}, ""))

}
