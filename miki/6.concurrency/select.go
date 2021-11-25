package main

import (
	"log"
	"time"
)

func main() {
	ch1, ch2 := make(chan int), make(chan int)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- 42
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- 24
	}()

	select {
	case v := <-ch1:
		log.Println("ch1", v)
	case v := <-ch2:
		log.Println("ch2", v)
	}
}
