package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"sync"
	"time"
)

func main() {
	start := time.Now()

	procs := 8
	runtime.GOMAXPROCS(procs)
	fmt.Println("num cpu:", runtime.NumCPU())
	fmt.Println("max procs:", runtime.GOMAXPROCS(procs))

	var wg sync.WaitGroup

	urls := []string{
		"https://hashnode.com/",
		"https://dev.to/",
		"https://stackoverflow.com/",
		"https://go.dev/",
		"https://medium.com/",
		"https://github.com/",
		"https://techcrunch.com/",
		"https://www.techrepublic.com/",
	}

	wg.Add(len(urls))
	for _, url := range urls {
		go check(url, &wg)
	}
	wg.Wait()

	fmt.Println("Exe time:", time.Since(start))
}

func check(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	_, err := http.Get(url)
	if err != nil {
		log.Println("\tERR->", err)
	}
	fmt.Println(url, "LIVE")
}
