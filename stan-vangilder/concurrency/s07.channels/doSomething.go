package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)

	go doSomething1(ch)
	go doSomething2(ch)

	fmt.Println(<-ch)
	fmt.Println(<-ch)

	close(ch)

	log.Println("\tFinished in", time.Since(start))
}

func doSomething1(ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "finished 1"
}

func doSomething2(ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "finished 2"
}
