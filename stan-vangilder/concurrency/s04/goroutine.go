package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	var wg sync.WaitGroup
	log.Println("--starting tasks--")
	wg.Add(1)
	go makeTea(&wg)
	wg.Add(1)
	go makeBed(&wg)
	wg.Wait()
	log.Println("--finished all tasks--")
	since := time.Since(start)
	log.Printf("took %v\n", since)
}

func makeTea(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(2 * time.Second)
	log.Println("made tea")
}

func makeBed(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(2 * time.Second)
	log.Println("made bed")
}
