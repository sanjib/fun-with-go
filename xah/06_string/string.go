package main

import "fmt"

func main() {
	x := "abc and ♥"
	fmt.Println(x)

	x = "he said, \"wow\""
	fmt.Println(x)

	x = `she said, "wow"
		 and then she said,
	"nice"`
	fmt.Println(x)
}
