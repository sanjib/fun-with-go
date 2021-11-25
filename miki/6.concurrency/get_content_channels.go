package main

import "log"

func getContentTypeForChannel(url string, ch chan string) {
	//resp, err := http.Get(url)
	//if err != nil {
	//	log.Println(err)
	//}
	//defer resp.Body.Close()
	//
	//ct := resp.Header.Get("Content-Type")
	//ch <- fmt.Sprint(url, ": ", ct)

	ch <- "he.."
}

func main() {
	urls := []string{
		"https://google.com",
		"https://golang.org",
		"https://go.dev",
	}
	ch := make(chan string)

	for _, url := range urls {
		//go func(url string) {
		//	err := getContentTypeForChannel(url, ch)
		//	if err != nil {
		//		log.Println(err)
		//	}
		//}(url)
		go getContentTypeForChannel(url, ch)
	}

	for range urls {
		log.Println(<-ch)
	}
}
