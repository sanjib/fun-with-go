package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	var wg sync.WaitGroup
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(n int, wg *sync.WaitGroup) {
			defer wg.Done()
			r := rand.Intn(10000)
			duration := time.Duration(r) * time.Millisecond
			log.Printf("%d. sleeping for %v\n", n, duration)
			time.Sleep(duration)
		}(i, &wg)
	}
	wg.Wait()
	log.Printf("--took %v--", time.Since(start))
}
