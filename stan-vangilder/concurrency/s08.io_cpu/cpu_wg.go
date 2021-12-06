package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	start := time.Now()

	var wg sync.WaitGroup
	p := 8
	runtime.GOMAXPROCS(p)
	fmt.Println("max procs:", runtime.GOMAXPROCS(p))

	countA := func(wg *sync.WaitGroup) {
		defer wg.Done()
		fmt.Println("AAA is starting")
		for i := 0; i < 3_000_000_000; i++ {
		}
		fmt.Println("AAA is done")
	}

	countB := func(wg *sync.WaitGroup) {
		defer wg.Done()
		fmt.Println("BBB is starting")
		for i := 0; i < 3_000_000_000; i++ {
		}
		fmt.Println("BBB is done")
	}

	countC := func(wg *sync.WaitGroup) {
		defer wg.Done()
		fmt.Println("CCC is starting")
		for i := 0; i < 3_000_000_000; i++ {
		}
		fmt.Println("CCC is done")
	}

	countD := func(wg *sync.WaitGroup) {
		defer wg.Done()
		fmt.Println("DDD is starting")
		for i := 0; i < 3_000_000_000; i++ {
		}
		fmt.Println("DDD is done")
	}

	countE := func(wg *sync.WaitGroup) {
		defer wg.Done()
		fmt.Println("EEE is starting")
		for i := 0; i < 3_000_000_000; i++ {
		}
		fmt.Println("EEE is done")
	}

	countF := func(wg *sync.WaitGroup) {
		defer wg.Done()
		fmt.Println("FFF is starting")
		for i := 0; i < 3_000_000_000; i++ {
		}
		fmt.Println("FFF is done")
	}

	countG := func(wg *sync.WaitGroup) {
		defer wg.Done()
		fmt.Println("GGG is starting")
		for i := 0; i < 3_000_000_000; i++ {
		}
		fmt.Println("GGG is done")
	}

	countH := func(wg *sync.WaitGroup) {
		defer wg.Done()
		fmt.Println("HHH is starting")
		for i := 0; i < 3_000_000_000; i++ {
		}
		fmt.Println("HHH is done")
	}

	wg.Add(8)
	go countA(&wg)
	go countB(&wg)
	go countC(&wg)
	go countD(&wg)
	go countE(&wg)
	go countF(&wg)
	go countG(&wg)
	go countH(&wg)

	wg.Wait()

	fmt.Println("-->", time.Since(start))
}
