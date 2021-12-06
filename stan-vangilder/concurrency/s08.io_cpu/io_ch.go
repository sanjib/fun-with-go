package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"time"
)

func main() {
	start := time.Now()

	procs := 8
	runtime.GOMAXPROCS(procs)
	fmt.Println("num cpu:", runtime.NumCPU())
	fmt.Println("max procs:", runtime.GOMAXPROCS(procs))

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

	ch := make(chan string, len(urls))
	for _, url := range urls {
		go checkWithChannel(url, ch)
	}

	for range urls {
		fmt.Println(<-ch)
	}

	close(ch)
	fmt.Println("Exe time:", time.Since(start))
}

func checkWithChannel(url string, ch chan string) {
	_, err := http.Get(url)
	if err != nil {
		log.Println("\tERR->", err)
	}
	ch <- fmt.Sprintf("%s is LIVE", url)
}
