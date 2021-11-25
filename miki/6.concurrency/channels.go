package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(time.Second * 3)
			ch <- i
		}
		close(ch)
	}()

	for v := range ch {
		fmt.Printf("got %d\n", v)
	}

}
