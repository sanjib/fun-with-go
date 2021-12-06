package main

import "fmt"

func main() {
	c := make(chan string, 3)
	c <- "hello "
	c <- "earth "
	c <- "from mars "

	fmt.Print(<-c)

	c <- "from venus "

	fmt.Print(<-c)
	fmt.Print(<-c)
	fmt.Print(<-c)

}
