package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)

	go func(c chan string) {
		for {
			time.Sleep(time.Second)
			c <- "sending every 1 sec"
		}
	}(c1)

	go func(c chan string) {
		for {
			time.Sleep(time.Second * 4)
			c <- "sending every 4 sec"
		}
	}(c2)

	go func(c chan string) {
		for {
			time.Sleep(time.Second * 10)
			c <- "We're done!"
		}
	}(c3)

	for {
		select {
		case msg := <-c1:
			fmt.Println(msg)
		case msg := <-c2:
			fmt.Println(msg + " something cool happened!")
		case msg := <-c3:
			fmt.Println(msg)
			os.Exit(1)
		}
	}
}
