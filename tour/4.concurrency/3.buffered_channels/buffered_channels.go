package main

import "fmt"

func main() {
	ch := make(chan int, 4)
	ch <- 100
	ch <- 120
	ch <- 110
	ch <- 130
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
