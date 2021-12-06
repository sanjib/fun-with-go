package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 100_000; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			ch := make(chan int)
			blocker := make(chan string)

			calcSquare := func(val int, ch chan int) {
				time.Sleep(time.Second * 3)
				ch <- val * val
			}

			report := func(val int, ch chan int, blocker chan string) {
				time.Sleep(time.Second * 1)
				fmt.Println(val, "square =", <-ch)
				blocker <- "fake channel to block"
			}

			go calcSquare(n, ch)
			go report(n, ch, blocker)

			<-blocker
		}(i)
	}
	wg.Wait()

}
