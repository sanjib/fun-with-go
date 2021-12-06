package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	start := time.Now()

	p := 8
	runtime.GOMAXPROCS(p)
	fmt.Println("max procs:", runtime.GOMAXPROCS(p))

	countA := func(ch chan string) {
		ch <- "AAA is starting"
		for i := 0; i < 3_000_000_000; i++ {
		}
		ch <- "AAA is done"
	}

	countB := func(ch chan string) {
		ch <- "BBB is starting"
		for i := 0; i < 3_000_000_000; i++ {
		}
		ch <- "BBB is done"
	}

	countC := func(ch chan string) {
		ch <- "CCC is starting"
		for i := 0; i < 3_000_000_000; i++ {
		}
		ch <- "CCC is done"
	}

	countD := func(ch chan string) {
		ch <- "DDD is starting"
		for i := 0; i < 3_000_000_000; i++ {
		}
		ch <- "DDD is done"
	}

	countE := func(ch chan string) {
		ch <- "EEE is starting"
		for i := 0; i < 3_000_000_000; i++ {
		}
		ch <- "EEE is done"
	}

	countF := func(ch chan string) {
		ch <- "FFF is starting"
		for i := 0; i < 3_000_000_000; i++ {
		}
		ch <- "FFF is done"
	}

	countG := func(ch chan string) {
		ch <- "GGG is starting"
		for i := 0; i < 3_000_000_000; i++ {
		}
		ch <- "GGG is done"
	}

	countH := func(ch chan string) {
		ch <- "HHH is starting"
		for i := 0; i < 3_000_000_000; i++ {
		}
		ch <- "HHH is done"
	}

	ch := make(chan string)
	go countA(ch)
	go countB(ch)
	go countC(ch)
	go countD(ch)
	go countE(ch)
	go countF(ch)
	go countG(ch)
	go countH(ch)

	for i := 0; i < 16; i++ {
		fmt.Println(<-ch)
	}

	fmt.Println("-->", time.Since(start))
}
