package main

import (
	"log"
	"net/http"
	"sync"
)

func getContentType(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	ct := resp.Header.Get("Content-Type")
	log.Println(url, "->", ct)
	return ct, nil
}

func main() {
	urls := []string{
		"https://google.com",
		"https://golang.org",
		"https://go.dev",
	}

	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			getContentType(url)
			wg.Done()
		}(url)
	}
	wg.Wait()
}
