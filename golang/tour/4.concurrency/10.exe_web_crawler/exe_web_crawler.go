package main

import (
	"fmt"
	"sync"
)

// Fetcher interface
type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

var didFetch = map[string]bool{}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, ch chan string) {
	var mu sync.Mutex
	defer close(ch)
	if depth <= 0 {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	mu.Lock()
	didFetch[url] = true
	mu.Unlock()

	ch <- fmt.Sprintf("found: %s %q", url, body)

	moreCh := make([]chan string, len(urls))
	for i, u := range urls {

		mu.Lock()
		_, ok := didFetch[u]
		mu.Unlock()

		if ok == false {
			moreCh[i] = make(chan string)
			go Crawl(u, depth-1, fetcher, moreCh[i])
			didFetch[u] = true

			for v := range moreCh[i] {
				ch <- v
			}
		}
	}

	return
}

func main() {
	ch := make(chan string)
	go Crawl("https://golang.org/", 4, fetcher, ch)

	for v := range ch {
		fmt.Println(v)
	}
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
